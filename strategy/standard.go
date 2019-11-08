// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package strategy

import (
	"context"
	"encoding/binary"
	"math/big"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/awishformore/cancoillotte/adaptor/gatekeeper"
	"github.com/awishformore/cancoillotte/adaptor/resolver"
	"github.com/awishformore/cancoillotte/adaptor/tournament"
	"github.com/awishformore/cancoillotte/adaptor/website"
	"github.com/awishformore/cancoillotte/adaptor/wizardguild"
	"github.com/awishformore/cancoillotte/model"
	"github.com/awishformore/cancoillotte/model/challenge"
	"github.com/awishformore/cancoillotte/model/duel"
	"github.com/awishformore/cancoillotte/model/outcome"
	"github.com/awishformore/cancoillotte/model/strategy"
	"github.com/awishformore/cancoillotte/model/wizard"
	"github.com/awishformore/cancoillotte/state/duels"
	"github.com/awishformore/cancoillotte/state/players"
	"github.com/awishformore/cancoillotte/state/secrets"
	"github.com/awishformore/cancoillotte/state/wizards"
	"github.com/awishformore/cancoillotte/state/world"
)

// Standard implements the different scenarios of our strategy.
type Standard struct {
	log    zerolog.Logger
	ctx    context.Context
	cfg    Config
	guild  *wizardguild.Binding
	gate   *gatekeeper.Binding
	tour   *tournament.Binding
	reso   *resolver.Binding
	wiz    *wizards.Repo
	duel   *duels.Repo
	sec    *secrets.Repo
	play   *players.Repo
	info   *world.Info
	web    *website.Client
	key    *keystore.Key
	sets   []model.Set
	height uint64
	phase  string
	window string
	sim    bool
}

// NewStandard initializes the strategy with dependencies.
func NewStandard(log zerolog.Logger, ctx context.Context, guild *wizardguild.Binding, gate *gatekeeper.Binding, tour *tournament.Binding, reso *resolver.Binding, wiz *wizards.Repo, duels *duels.Repo, sec *secrets.Repo, play *players.Repo, info *world.Info, web *website.Client, key *keystore.Key, cfg Config, sim bool) (*Standard, error) {

	height, err := tour.Height()
	if err != nil {
		return nil, errors.Wrap(err, "could not get height")
	}

	phase, err := info.Phase()
	if err != nil {
		return nil, errors.Wrap(err, "could not get phase")
	}

	window, err := info.Window()
	if err != nil {
		return nil, errors.Wrap(err, "could not get window")
	}

	sets := make([]model.Set, 0, 3*3*3*3*3)
	for a1 := byte(2); a1 <= 4; a1++ {
		for a2 := byte(2); a2 <= 4; a2++ {
			for a3 := byte(2); a3 <= 4; a3++ {
				for a4 := byte(2); a4 <= 4; a4++ {
					for a5 := byte(2); a5 <= 4; a5++ {
						set := model.Set{a1, a2, a3, a4, a5}
						valid, err := reso.Valid(set)
						if err != nil {
							return nil, errors.Wrap(err, "could net check set validity")
						}
						if !valid {
							return nil, errors.Errorf("generated invalid set (%x)", set)
						}
						sets = append(sets, set)
					}
				}
			}
		}
	}

	log.Info().
		Str("address", info.Address()).
		Uint64("height", height).
		Str("phase", phase).
		Str("window", window).
		Int("sets", len(sets)).
		Uint64("reserve", cfg.Reserve).
		Uint64("extra", cfg.Extra).
		Uint("neutral", cfg.Ratio.Neutral).
		Uint("fire", cfg.Ratio.Fire).
		Uint("water", cfg.Ratio.Water).
		Uint("wind", cfg.Ratio.Wind).
		Msg("standard strategy initialized")

	s := &Standard{
		log:    log,
		ctx:    ctx,
		cfg:    cfg,
		guild:  guild,
		gate:   gate,
		tour:   tour,
		reso:   reso,
		wiz:    wiz,
		duel:   duels,
		sec:    sec,
		play:   play,
		info:   info,
		web:    web,
		key:    key,
		sets:   sets,
		height: height,
		phase:  phase,
		window: window,
		sim:    sim,
	}

	return s, nil
}

// OnBlock is called when we receive a block.
func (s *Standard) OnBlock(hash string, height uint64) error {

	if height < s.height {
		s.log.Warn().
			Uint64("old_height", s.height).
			Uint64("new_height", height).
			Msg("block rollback")
	}

	if height == s.height {
		s.log.Debug().
			Uint64("height", height).
			Msg("block fork")
	}

	s.height = height

	phase, err := s.info.Phase()
	if err != nil {
		return errors.Wrap(err, "could not retrieve phase")
	}

	if phase != s.phase {
		s.log.Info().
			Str("old_phase", string(s.phase)).
			Str("new_phase", string(phase)).
			Msg("tournament phase changed")
		s.phase = phase
	}

	window, err := s.info.Window()
	if err != nil {
		return errors.Wrap(err, "could not retrieve window")
	}

	if window != s.window {
		s.log.Info().
			Str("old_window", string(s.window)).
			Str("new_window", string(window)).
			Msg("tournament window changed")
		s.window = window
	}

	s.tick(height)

	return nil
}

// OnCreation is called when a new wizard is created.
func (s *Standard) OnCreation(wizID uint64, aff uint8, pow uint64) error {

	owner, err := s.guild.Owner(wizID)
	if err != nil {
		return errors.Wrap(err, "could not get owner")
	}

	player, err := s.web.Player(wizID)
	if err != nil {
		return errors.Wrap(err, "could not get player")
	}

	s.wiz.Update(wizID, func(wiz *model.Wizard) {
		wiz.ID = wizID
		wiz.Owner = owner
		wiz.Player = player
		wiz.Aff = aff
		wiz.Pow = pow
		wiz.Ready = true
	})

	log := s.log.With().
		Uint64("wiz", wizID).
		Str("owner", owner).
		Uint8("aff", aff).
		Uint64("pow", pow).
		Logger()

	address := s.info.Address()
	if address == owner {
		log.Info().Msg("team wizard created")
	}

	return nil
}

// OnTransfer is called when a wizard is transferred to a new owner.
func (s *Standard) OnTransfer(wizID uint64, sender string, owner string) error {

	player, err := s.web.Player(wizID)
	if err != nil {
		return errors.Wrap(err, "could not get player")
	}

	s.wiz.Update(wizID, func(wiz *model.Wizard) {
		wiz.Owner = owner
		wiz.Player = player
	})

	log := s.log.With().
		Uint64("wiz", wizID).
		Str("sender", sender).
		Str("owner", owner).
		Logger()

	address := s.info.Address()
	if address == sender {
		log.Info().Msg("outgoing wizard transfer completed")
	}
	if address == owner {
		log.Info().Msg("incoming wizard transfer completed")
	}

	return nil
}

// OnChallenge is called when a wizard issues a challenge.
func (s *Standard) OnChallenge(wizID1 uint64, wizID2 uint64, nonce1 uint64, nonce2 uint64, commit1 model.Hash) error {

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	log := s.log.With().
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz1.Owner {
		log.Info().Msg("outgoing on-chain challenge issued")
	}
	if address == wiz2.Owner {
		log.Info().Msg("incoming on-chain challenge issued")
	}

	return nil
}

// OnWithdrawal is called when a wizard cancels his challenge.
func (s *Standard) OnWithdrawal(wizID uint64) error {

	s.wiz.Update(wizID, func(wiz *model.Wizard) {
		wiz.Ready = true
	})

	wiz := s.wiz.Get(wizID)

	log := s.log.With().
		Str("owner", wiz.Owner).
		Uint64("wizID", wizID).
		Logger()

	address := s.info.Address()
	if address == wiz.Owner {
		log.Info().Msg("outgoing on-chain challenge withdrawn")
	}

	return nil
}

// OnFight is called when a duel is accepted.
func (s *Standard) OnFight(duelID model.Hash, wizID1 uint64, wizID2 uint64, timeout uint64, contest bool, commit1 model.Hash, commit2 model.Hash) error {

	s.wiz.Update(wizID1, func(wiz *model.Wizard) {
		wiz.Ready = false
	})

	s.wiz.Update(wizID2, func(wiz *model.Wizard) {
		wiz.Ready = false
	})

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	s.info.Action(wiz1.Player, s.height)
	s.info.Action(wiz2.Player, s.height)

	err := s.duel.Update(duelID, func(duel *model.Duel) {
		duel.ID = duelID
		duel.Owner1 = wiz1.Owner
		duel.Owner2 = wiz2.Owner
		duel.Player1 = wiz1.Player
		duel.Player2 = wiz2.Player
		duel.Wiz1 = wizID1
		duel.Wiz2 = wizID2
		duel.Aff1 = wiz1.Aff
		duel.Aff2 = wiz2.Aff
		duel.Start = s.height
		duel.StartPow1 = wiz1.Pow
		duel.StartPow2 = wiz2.Pow
		duel.Deadline = timeout
	})
	if err != nil {
		return errors.Wrap(err, "could not update duel")
	}

	log := s.log.With().
		Hex("duel", duelID[:]).
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	address := s.info.Address()
	if wiz1.Owner == address {
		log.Info().Msg("home wizard challenge accepted")
	}
	if wiz2.Owner == address {
		log.Info().Msg("away wizard challenge accepted")
	}

	return nil
}

// OnReveal is called when movesets are revealed by one wizard.
func (s *Standard) OnReveal(duelID model.Hash, wizID1 uint64, wizID2 uint64) error {

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	log := s.log.With().
		Hex("duel", duelID[:]).
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz1.Owner {
		log.Info().Msg("home wizard reveal submitted")
	}
	if address == wiz2.Owner {
		log.Info().Msg("away wizard reveal submitted")
	}

	return nil
}

// OnResult is called when a duel has ended.
func (s *Standard) OnResult(duelID model.Hash, wizID1 uint64, wizID2 uint64, set1 model.Set, set2 model.Set, pow1 uint64, pow2 uint64) error {

	s.wiz.Update(wizID1, func(wiz *model.Wizard) {
		wiz.Pow = pow1
		wiz.Ready = true
	})

	s.wiz.Update(wizID2, func(wiz *model.Wizard) {
		wiz.Pow = pow2
		wiz.Ready = true
	})

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	s.info.Action(wiz1.Player, s.height)
	s.info.Action(wiz2.Player, s.height)

	err := s.duel.Update(duelID, func(duel *model.Duel) {
		duel.End = s.height
		duel.Set1 = set1
		duel.Set2 = set2
		duel.EndPow1 = pow1
		duel.EndPow2 = pow2
	})
	if err != nil {
		return errors.Wrap(err, "could not update duel")
	}

	log := s.log.With().
		Hex("duel", duelID[:]).
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Hex("set1", set1[:]).
		Hex("set2", set2[:]).
		Logger()

	address := s.info.Address()
	if address == wiz1.Owner {
		log.Info().Msg("home wizard duel concluded")
	}
	if address == wiz2.Owner {
		log.Info().Msg("away wizard duel concluded")
	}

	return nil
}

// OnTimeout is called when a duel times out.
func (s *Standard) OnTimeout(duelID model.Hash, wizID1 uint64, wizID2 uint64, pow1 uint64, pow2 uint64) error {

	s.wiz.Update(wizID1, func(wiz *model.Wizard) {
		wiz.Pow = pow1
		wiz.Ready = true
	})

	s.wiz.Update(wizID2, func(wiz *model.Wizard) {
		wiz.Pow = pow2
		wiz.Ready = true
	})

	err := s.duel.Update(duelID, func(duel *model.Duel) {
		duel.End = s.height
		duel.EndPow1 = pow1
		duel.EndPow2 = pow2
		duel.Timeout = true
	})
	if err != nil {
		return errors.Wrap(err, "could not update duel")
	}

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	log := s.log.With().
		Hex("duel", duelID[:]).
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz1.Owner {
		log.Info().Msg("home wizard duel timeout")
	}
	if address == wiz2.Owner {
		log.Info().Msg("away wizard duel timeout")
	}

	return nil
}

// OnElimination is called when a wizard is eliminated.
func (s *Standard) OnElimination(wizID uint64) error {

	s.wiz.Update(wizID, func(wiz *model.Wizard) {
		wiz.Pow = 0
	})

	wiz := s.wiz.Get(wizID)

	log := s.log.With().
		Str("owner", wiz.Owner).
		Uint64("wiz", wizID).
		Uint8("aff", wiz.Aff).
		Uint64("pow", wiz.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz.Owner {
		log.Info().Msg("team wizard eliminated")
	}

	return nil
}

// OnAttempt is called when a wizard attempts ascension.
func (s *Standard) OnAttempt(wizID uint64) error {

	wiz := s.wiz.Get(wizID)

	s.info.Action(wiz.Player, s.height)

	log := s.log.With().
		Str("owner", wiz.Owner).
		Uint64("wiz", wizID).
		Uint8("aff", wiz.Aff).
		Uint64("pow", wiz.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz.Owner {
		log.Info().Msg("team wizard attempting ascension")
	}
	if address != wiz.Owner {
		log.Info().Msg("enemy wizard attempting ascension")
	}

	return nil
}

// OnPair is called when an ascension attempt is paired.
func (s *Standard) OnPair(wizID1 uint64, wizID2 uint64) error {

	s.wiz.Update(wizID1, func(wiz *model.Wizard) {
		wiz.Ready = false
	})

	s.wiz.Update(wizID2, func(wiz *model.Wizard) {
		wiz.Ready = false
	})

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	s.info.Action(wiz2.Player, s.height)

	log := s.log.With().
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz1.Owner {
		log.Info().Msg("team wizard ascension thwarted")
	}
	if address == wiz2.Owner {
		log.Info().Msg("team wizard thwarted ascension")
	}
	if address != wiz1.Owner && address != wiz2.Owner {
		log.Info().Msg("enemy wizard ascension thwarted")
	}

	return nil
}

// OnContest is called when an ascending wizard is challenged.
func (s *Standard) OnContest(wizID1 uint64, wizID2 uint64, commit2 model.Hash) error {

	s.wiz.Update(wizID1, func(wiz *model.Wizard) {
		wiz.Ready = false
	})

	s.wiz.Update(wizID2, func(wiz *model.Wizard) {
		wiz.Ready = false
	})

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	s.info.Action(wiz2.Player, s.height)

	log := s.log.With().
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz1.Owner {
		log.Info().Msg("team wizard ascension challenged")
	}
	if address == wiz2.Owner {
		log.Info().Msg("team wizard challenged ascension")
	}
	if address != wiz1.Owner && address != wiz2.Owner {
		log.Info().Msg("enemy wizard ascension challenged")
	}

	return nil
}

// OnSuccess is called when an ascension attempt succeeded.
func (s *Standard) OnSuccess(wizID uint64, pow uint64) error {

	s.wiz.Update(wizID, func(wiz *model.Wizard) {
		wiz.Pow = pow
		wiz.Ready = true
	})

	wiz := s.wiz.Get(wizID)

	log := s.log.With().
		Str("owner", wiz.Owner).
		Uint64("wiz", wizID).
		Uint8("aff", wiz.Aff).
		Uint64("pow", wiz.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz.Owner {
		log.Info().Msg("team wizard ascension succeeded")
	}
	if address != wiz.Owner {
		log.Info().Msg("enemy wizard ascension succeeded")
	}

	return nil
}

// OnRevival is called when a wizard is revived.
func (s *Standard) OnRevival(wizID uint64, pow uint64) error {

	s.wiz.Update(wizID, func(wiz *model.Wizard) {
		wiz.Pow = pow
		wiz.Ready = true
	})

	wiz := s.wiz.Get(wizID)

	log := s.log.With().
		Str("owner", wiz.Owner).
		Uint64("wiz", wizID).
		Uint8("aff", wiz.Aff).
		Uint64("pow", wiz.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz.Owner {
		log.Info().Msg("team wizard revived")
	}

	return nil
}

// OnWipe is called when power is completely transfered.
func (s *Standard) OnWipe(wizID1 uint64, wizID2 uint64, pow uint64, reason uint8) error {

	s.wiz.Update(wizID1, func(wiz *model.Wizard) {
		wiz.Pow = 0
	})

	s.wiz.Update(wizID2, func(wiz *model.Wizard) {
		wiz.Pow = wiz.Pow + pow
	})

	wiz1 := s.wiz.Get(wizID1)
	wiz2 := s.wiz.Get(wizID2)

	log := s.log.With().
		Str("owner1", wiz1.Owner).
		Str("owner2", wiz2.Owner).
		Uint64("wiz1", wizID1).
		Uint64("wiz2", wizID2).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	address := s.info.Address()
	if address == wiz1.Owner {
		log.Info().Msg("enemy wiped team wizard")
	}
	if address == wiz2.Owner {
		log.Info().Msg("team wizard wiped enemy")
	}

	return nil
}

// OnClaim is called when the price is claimed.
func (s *Standard) OnClaim(wizID uint64, prize uint64) error {

	wiz := s.wiz.Get(wizID)

	log := s.log.With().
		Str("owner", wiz.Owner).
		Uint64("wiz", wizID).
		Uint8("aff", wiz.Aff).
		Uint64("pow", wiz.Pow).
		Uint64("prize", prize).
		Logger()

	address := s.info.Address()
	if address == wiz.Owner {
		log.Info().Msg("team wizard claimed prize")
	}
	if address != wiz.Owner {
		log.Info().Msg("enemy wizard claimed prize")
	}

	return nil
}

// tick will execute once a block to execute actions.
func (s *Standard) tick(height uint64) {

	log := s.log.With().
		Uint64("height", height).
		Logger()

	challenges, err := s.web.Challenges()
	if err != nil {
		log.Error().Err(err).Msg("could not retrieve challenges")
		return
	}

	ctx, cancel := context.WithTimeout(s.ctx, 13*time.Second)
	defer cancel()

	// process the different parts of the strategy
	s.timeout(ctx, height)
	s.summon(ctx)
	s.ascend(ctx)
	s.commit(ctx, challenges)
	s.reveal(ctx, challenges)
	s.issue(ctx, challenges)
}

// timeout will send the timeout transaction for timed out events.
func (s *Standard) timeout(ctx context.Context, height uint64) {

	log := s.log.With().Str("tick", "timeout").Logger()

	// check if we have timed out duels
	expired := s.duel.Select().
		Filter(duel.Set1Zero).
		Filter(duel.Set2Zero).
		Filter(duel.Not(duel.Timeout)).
		Filter(duel.Expired(height))
	if len(expired) == 0 {
		log.Debug().Msg("no timed out duels")
		return
	}

	// if we do, submit the solution
	for _, d := range expired {

		log := log.With().Hex("duel", d.ID[:]).
			Uint64("wiz1", d.Wiz1).
			Uint64("wiz2", d.Wiz2).
			Logger()

		address := s.info.Address()
		if d.Owner1 != address && d.Owner2 != address {
			log.Debug().Msg("ignoring timeout from other players")
			continue
		}

		// time out the duel
		hash, err := s.tour.Timeout(d.Wiz1, d.Wiz2, s.cfg.Extra)
		if err != nil {
			log.Error().Err(err).Msg("could not submit duel timeout")
			continue
		}

		log.Info().Str("hash", hash).Msg("duel timeout submitted")
	}

}

// summon will check if we can and should summon wizards and do so.
func (s *Standard) summon(ctx context.Context) {

	log := s.log.With().Str("tick", "summon").Logger()

	// check if we are in the right phase to summon
	if s.phase != world.Admission && s.phase != world.Revival {
		log.Debug().Str("phase", s.phase).Msg("wrong phase for summoning wizards")
		return
	}

	// get the current player reserve and balance
	reserve := s.cfg.Reserve
	balance, err := s.info.Balance()
	if err != nil {
		log.Error().Err(err).Msg("could not get balance")
		return
	}

	log = log.With().
		Float64("balance", float64(balance)/1e18).
		Float64("reserve", float64(reserve)/1e18).
		Logger()

	// get the current costs of wizards
	basic, elemental, err := s.gate.Costs()
	if err != nil {
		log.Error().Err(err).Msg("could not check costs")
		return
	}

	// configure the costs by affinity
	costs := map[uint8]uint64{
		wizard.Neutral: basic,
		wizard.Fire:    elemental,
		wizard.Water:   elemental,
		wizard.Wind:    elemental,
	}

	// calculate the total configured probabilities
	total := float64(s.cfg.Ratio.Neutral + s.cfg.Ratio.Fire + s.cfg.Ratio.Water + s.cfg.Ratio.Wind)
	fire := float64(s.cfg.Ratio.Fire) / total
	water := float64(s.cfg.Ratio.Water) / total
	wind := float64(s.cfg.Ratio.Wind) / total

	// pick the affinity
	var aff uint8
	dice := rand.Float64()
	switch {
	case dice < fire:
		aff = wizard.Fire
	case dice < fire+water:
		aff = wizard.Water
	case dice < fire+water+wind:
		aff = wizard.Wind
	default:
		aff = wizard.Neutral
	}

	// deduce the cost
	cost := costs[aff]

	log = log.With().
		Uint8("aff", aff).
		Float64("cost", float64(cost)/1e18).
		Logger()

	// check if we can afford the summoning
	buffer := int64(balance) - int64(reserve)
	if buffer < int64(cost) {
		log.Debug().Msg("no pending summons")
		return
	}

	// summon the wizard
	hash, err := s.gate.Conjure(aff, cost)
	if err != nil {
		log.Error().Err(err).Msg("could not summon wizard")
		return
	}

	log.Info().Str("hash", hash).Msg("wizard summoning submitted")
}

// reveal will check if there are any duels we need to reveal the commit for.
func (s *Standard) reveal(ctx context.Context, challenges model.ChallengeList) {

	log := s.log.With().Str("tick", "reveal").Logger()

	if s.window != world.Fight && s.window != world.Resolution {
		log.Debug().Str("window", s.window).Msg("wrong window for revealing challenges")
		return
	}

	address := s.info.Address()

	// filter for challenges where both sides can reveal
	pending := challenges.
		Filter(challenge.Status(challenge.RevealReady))

	// filter for challenges where we are home and away revealed
	pending = append(pending, challenges.
		Filter(challenge.Status(challenge.RevealAway)).
		Filter(challenge.Owner1(address))...)

	// filter for challenges where we are away and home revealed
	pending = append(pending, challenges.
		Filter(challenge.Status(challenge.RevealHome)).
		Filter(challenge.Owner2(address))...)

	// skip if there are no pending reveals
	if len(pending) == 0 {
		log.Debug().Msg("no pending reveals")
		return
	}

	for _, cha := range pending {

	RevealLoop:
		select {
		case <-ctx.Done():
			log.Debug().Msg("leaving reveal loop early")
			break RevealLoop
		default:
		}

		log := log.With().
			Str("challenge", cha.ID).
			Uint64("wiz1", cha.Wiz1).
			Uint64("wiz2", cha.Wiz2).
			Str("status", cha.Status).
			Logger()

		// reveal if we are the home wizard
		if cha.Owner1 == address {

			// retrieve the secret associated with the home side
			secret, err := s.sec.Recover(cha.ID, cha.Wiz1)
			if err != nil {
				log.Error().Err(err).Msg("could not recover secret")
				continue
			}

			// reveal the salt and moves
			key := secret.Key(cha.Wiz1, cha.Wiz2)
			log = log.With().
				Hex("set", secret.Set[:]).
				Hex("salt", secret.Salt[:]).
				Hex("commit", secret.Commit[:]).
				Str("key", key).
				Logger()
			_, err = s.web.Reveal(cha.Wiz1, cha.ID, secret.Salt, secret.Set)
			if err != nil {
				log.Error().Err(err).Msg("could not submit home reveal")
				continue
			}

			log.Info().Msg("home reveal submitted")
		}

		// reveal if we are the away wizard
		if cha.Owner2 == address {

			// retrieve the secret associated with the away side
			secret, err := s.sec.Recover(cha.ID, cha.Wiz2)
			if err != nil {
				log.Error().Err(err).Msg("could not recover secret")
				continue
			}

			// reveal the salt and moves
			key := secret.Key(cha.Wiz1, cha.Wiz2)
			log = log.With().
				Hex("set", secret.Set[:]).
				Hex("salt", secret.Salt[:]).
				Hex("commit", secret.Commit[:]).
				Str("key", key).
				Logger()
			_, err = s.web.Reveal(cha.Wiz2, cha.ID, secret.Salt, secret.Set)
			if err != nil {
				log.Error().Err(err).Msg("could not submit away reveal")
				continue
			}

			log.Info().Msg("away reveal submitted")
		}
	}
}

// ascend will check if we can attempt an ascend or challenge ascending.
func (s *Standard) ascend(ctx context.Context) {

}

// commit will check if there are any pending challenges to commit.
func (s *Standard) commit(ctx context.Context, challenges model.ChallengeList) {

	log := s.log.With().Str("tick", "commit").Logger()

	if s.window != world.Fight {
		log.Debug().Str("window", s.window).Msg("wrong window for accepting challenges")
		return
	}

	var pending model.ChallengeList

	// get all pending challenges that we sent
	outgoing := challenges.
		Filter(challenge.Owner1(s.info.Address())).
		Filter(challenge.Status(challenge.CommitAway))
	if len(outgoing) != 0 {
		log.Debug().Int("outgoing", len(outgoing)).Msg("found pending outgoing challenges")
		pending = append(pending, outgoing...)
	}

	// get all pending challenges that we received
	incoming := challenges.
		Filter(challenge.Owner2(s.info.Address())).
		Filter(challenge.Status(challenge.Sent))
	if len(incoming) != 0 {
		log.Debug().Int("incoming", len(incoming)).Msg("found pending incoming challenges")
		pending = append(pending, incoming...)
	}

	if len(pending) == 0 {
		log.Debug().Msg("no pending commits")
		return
	}

	address := s.info.Address()

	// send our commit for all pending challenges
CommitLoop:
	for _, cha := range pending {

		select {
		case <-ctx.Done():
			log.Debug().Msg("leaving commit loop early")
			break CommitLoop
		default:
		}

		log := log.With().Str("challenge", cha.ID).Logger()

		// get the first wizard from the contract for nonce
		wiz1, err := s.tour.Wizard(cha.Wiz1)
		if err != nil {
			log.Error().Err(err).Msg("could not get home wizard")
			continue
		}

		// get the second wizard from the contract for nonce
		wiz2, err := s.tour.Wizard(cha.Wiz2)
		if err != nil {
			log.Error().Err(err).Msg("could not get away wizard")
			continue
		}

		// generate commit for home side
		if cha.Owner1 == address {

			// calculate best strategy and pick a random set
			strat := s.pickStrategy(s.duel.Select(), wiz1, wiz2)
			set := strat.Sets[rand.Intn(len(strat.Sets))]

			// store the set to get the commit
			secret := generateSecret(cha.ID, set, wiz1.ID)
			err = s.sec.Backup(secret)
			if err != nil {
				log.Error().Err(err).Msg("could not backup secret")
				continue
			}

			log = log.With().
				Uint64("wiz1", cha.Wiz1).
				Uint64("wiz2", cha.Wiz2).
				Uint32("nonce1", wiz1.Nonce).
				Uint32("nonce2", wiz2.Nonce).
				Hex("set", secret.Set[:]).
				Hex("salt", secret.Salt[:]).
				Hex("commit", secret.Commit[:]).
				Str("status", cha.Status).
				Logger()

			// get the signature
			sig, err := s.getSignature(cha.Wiz1, cha.Wiz2, wiz1.Nonce, wiz2.Nonce, secret.Commit)
			if err != nil {
				log.Error().Err(err).Msg("could not get signature")
				continue
			}

			log = log.With().Hex("sig", sig).Logger()

			// send the commit to the web API
			_, err = s.web.Commit(cha.Wiz1, s.info.Address(), cha.ID, secret.Commit, sig)
			if err != nil {
				log.Error().Err(err).Msg("could not submit commit")
				continue
			}

			log.Info().Msg("home commit submitted")
		}

		// generate commit for away side
		if cha.Owner2 == address {

			// calculate best strategy and pick a random set
			strat := s.pickStrategy(s.duel.Select(), wiz2, wiz1)
			set := strat.Sets[rand.Intn(len(strat.Sets))]

			// store the set to get the commit
			secret := generateSecret(cha.ID, set, wiz2.ID)
			err = s.sec.Backup(secret)
			if err != nil {
				log.Error().Err(err).Msg("could not backup secret")
				continue
			}

			log = log.With().
				Uint64("wiz1", cha.Wiz1).
				Uint64("wiz2", cha.Wiz2).
				Uint32("nonce1", wiz1.Nonce).
				Uint32("nonce2", wiz2.Nonce).
				Hex("set", secret.Set[:]).
				Hex("salt", secret.Salt[:]).
				Hex("commit", secret.Commit[:]).
				Str("status", cha.Status).
				Logger()

			// get the signature
			sig, err := s.getSignature(cha.Wiz1, cha.Wiz2, wiz1.Nonce, wiz2.Nonce, secret.Commit)
			if err != nil {
				log.Error().Err(err).Msg("could not get signature")
				continue
			}

			log = log.With().Hex("sig", sig).Logger()

			// send the commit to the web API
			_, err = s.web.Commit(cha.Wiz1, s.info.Address(), cha.ID, secret.Commit, sig)
			if err != nil {
				log.Error().Err(err).Msg("could not submit commit")
				continue
			}

			log.Info().Msg("away commit submitted")
		}
	}
}

// issue will check if we can issue a challenge for any wizard.
func (s *Standard) issue(ctx context.Context, challenges model.ChallengeList) {

	log := s.log.With().Str("tick", "issue").Logger()

	if !s.sim && (s.window != world.Fight) {
		log.Debug().Str("window", s.window).Msg("wrong window for issuing challenges")
		return
	}

	// get challenges sorted by us on our the home side
	swapped := challenges.Swap(challenge.Owner2(s.info.Address()))

	// get all of our wizards that are already fighting
	busyHome := swapped.
		Filter(challenge.Owner1(s.info.Address())).
		Filter(challenge.Not(challenge.Status(challenge.Sent))).
		Filter(challenge.Not(challenge.Status(challenge.Done))).
		Filter(challenge.Not(challenge.Status(challenge.Rejected))).
		HomeIDs()

	busyAway := swapped.
		Filter(challenge.Not(challenge.Owner2(s.info.Address()))).
		Filter(challenge.Not(challenge.Status(challenge.Sent))).
		Filter(challenge.Not(challenge.Status(challenge.Done))).
		Filter(challenge.Not(challenge.Status(challenge.Rejected))).
		AwayIDs()

	// get all team wizards that are not busy
	team := s.wiz.Select().
		Filter(wizard.Owner(s.info.Address())).
		Filter(wizard.Available).
		Filter(wizard.Not(wizard.ID(busyHome...))).
		Random()
	if len(team) == 0 {
		log.Debug().Msg("no available team wizards")
		return
	}

	// get all enemy wizards that are active and available
	enemy := s.wiz.Select().
		Filter(wizard.Not(wizard.Owner(s.info.Address()))).
		Filter(wizard.Available).
		Filter(wizard.Active(s.info.Active, s.height))
	if len(enemy) == 0 {
		log.Debug().Msg("no available enemy wizards")
		return
	}

IssueLoop:
	for _, wiz1 := range team {

		select {
		case <-ctx.Done():
			log.Debug().Msg("leaving issue loop early")
			break IssueLoop
		default:
		}

		log := log.With().
			Uint64("wiz1", wiz1.ID).
			Uint8("aff1", wiz1.Aff).
			Uint64("pow1", wiz1.Pow).
			Logger()

		// get one random enemy wizard in our power range to challenge
		match := enemy.
			Filter(wizard.Not(wizard.ID(busyAway...))).
			Filter(wizard.Above(wiz1.Pow / 2)).
			Filter(wizard.Below(wiz1.Pow * 2)).
			Sample(1)
		if len(match) == 0 {
			log.Debug().Msg("no available wizard matches")
			continue
		}
		wiz2 := match[0]

		log = log.With().
			Uint64("wiz2", wiz2.ID).
			Uint8("aff2", wiz2.Aff).
			Uint64("pow2", wiz2.Pow).
			Logger()

		// if we are in simulation mode, just print the challenge
		if s.sim {
			duels := s.duel.Select()
			strat := s.pickStrategy(duels, wiz1, wiz2)
			set := strat.Sets[rand.Intn(len(strat.Sets))]
			log.Debug().
				Float64("score", strat.Score).
				Hex("set", set[:]).
				Msg("challenge issue simulated")
			continue
		}

		// submit challenge to get the challenge ID
		chaID, err := s.web.Issue(wiz1.ID, wiz2.ID)

		// we don't want to keep challenging this tick if we reached the limit
		if isTooManyErr(err) {
			log.Debug().Msg("challenge limit reached")
			break IssueLoop
		}

		// we don't want to keep challenging wizards that are busy
		if isInDuelErr(err) {
			_ = s.wiz.Update(wiz2.ID, func(wiz *model.Wizard) {
				wiz.Ready = false
			})
			continue
		}

		// otherwise, just log the error and keep going
		if err != nil {
			log.Error().Err(err).Msg("could not issue challenge")
			continue
		}

		// update wizard as busy
		busyAway = append(busyAway, wiz2.ID)

		log.Info().Str("challenge", chaID).Msg("duel challenge submitted")
	}
}

func (s *Standard) pickStrategy(duels model.DuelList, wiz1 model.Wizard, wiz2 model.Wizard) model.Strategy {

	log := log.With().
		Uint64("wiz1", wiz1.ID).
		Uint64("wiz2", wiz2.ID).
		Uint8("aff1", wiz1.Aff).
		Uint8("aff2", wiz2.Aff).
		Uint64("pow1", wiz1.Pow).
		Uint64("pow2", wiz2.Pow).
		Logger()

	// each strategy will have five sets
	var strategies model.StrategyList

	// sort duels from oldest to newest
	duels = duels.Sort(duel.ByStartAsc)

	// use all possible sets as our possible moves
	sets1 := s.sets

	// generate a stupid strategy based on just testing what works best against
	// this wizard, by testing each set against each set
	sets2 := s.sets
	strat := s.simulateStrategy(sets1, sets2, wiz1, wiz2)
	strategies = append(strategies, strat)

	// now we want to be a bit more fancy; we have different criteria related to
	// the both wizards that we can use to filter the training set; with each
	// selected training set, we run all possible movesets against them; note
	// that this includes level 2 approaches, where we just use our own filters
	// as selection, thus modeling what an opponent adjusting to us would do
	for homePlayer := 0; homePlayer <= 1; homePlayer++ {
		for awayPlayer := 0; awayPlayer <= 1; awayPlayer++ {
			for homeAffinity := 0; homeAffinity <= 1; homeAffinity++ {
				for awayAffinity := 0; awayAffinity <= 1; awayAffinity++ {
					for homeWizard := 0; homeWizard <= 1; homeWizard++ {
						for awayWizard := 0; awayWizard <= 1; awayWizard++ {
							for powerRange := 0; powerRange <= 3; powerRange++ {
								selection := duels.
									Filter(duel.Not(duel.Timeout)).
									Swap(duel.Player1(wiz2.Player))
								if homePlayer == 1 {
									selection = selection.Filter(duel.Player1(wiz1.Player))
								}
								if awayPlayer == 1 {
									selection = selection.Filter(duel.Player2(wiz2.Player))
								}
								if homeAffinity == 1 {
									selection = selection.Filter(duel.Aff1(wiz1.Aff))
								}
								if awayAffinity == 1 {
									selection = selection.Filter(duel.Aff2(wiz2.Aff))
								}
								if homeWizard == 1 {
									selection = selection.Filter(duel.Wiz1(wiz1.ID))
								}
								if awayWizard == 1 {
									selection = selection.Filter(duel.Wiz2(wiz2.ID))
								}
								if powerRange == 1 {
									selection = selection.Filter(duel.Within(1.2))
								}
								if powerRange == 2 {
									selection = selection.Filter(duel.Weaker(1.2))
								}
								if powerRange == 3 {
									selection = selection.Filter(duel.Stronger(1.2))
								}
								if len(selection) > 243 {
									selection = selection[:243]
								}
								if len(selection) >= 9 {
									sets2 := selection.Sets()
									strat := s.simulateStrategy(sets1, sets2, wiz1, wiz2)
									strategies = append(strategies, strat)
									log.Debug().
										Int("home_player", homePlayer).
										Int("away_player", awayPlayer).
										Int("home_aff", homeAffinity).
										Int("away_aff", awayAffinity).
										Int("home_wiz", homeWizard).
										Int("away_wiz", awayWizard).
										Int("selection", len(selection)).
										Float64("score", strat.Score).
										Msg("simulated strategy")
								}
							}
						}
					}
				}
			}
		}
	}

	// find the best strategy of those tested
	strategies = strategies.Random().Sort(strategy.Reverse(strategy.ScoreAsc))

	return strategies[0]
}

func (s *Standard) simulateStrategy(sets1 []model.Set, sets2 []model.Set, wiz1 model.Wizard, wiz2 model.Wizard) model.Strategy {

	// we first simulate an outcome for each possible set on our side versus
	// each set we selected for our opponent; we keep a moving average as the
	// basis for scoring the outcomes for this strategy, while we keep the
	// variance to evuate the strategy against other strategies
	outcomes := make(model.OutcomeList, 0, len(sets1))
	for _, set1 := range sets1 {

		// declare the variables to told calculations
		var average float64                      // average (needed for variance)
		var variance float64                     // variance (squared standard deviation)
		var ema float64                          // exponential moving average
		deltas := make([]float64, 0, len(sets1)) // all results (needed for variance)

		// we need to set the initial values with the first set
		set2 := sets2[0]
		delta := float64(resolveDuel(set1, set2, wiz1.Aff, wiz2.Aff, wiz1.Pow, wiz2.Pow))
		deltas = append(deltas, delta)
		average = delta
		ema = delta

		// then we go through the rest, accumulating results and calculating the EMA
		for _, set2 := range sets2 {
			delta := float64(resolveDuel(set1, set2, wiz1.Aff, wiz2.Aff, wiz1.Pow, wiz2.Pow))
			deltas = append(deltas, delta)
			average = average + delta
			ema = ema + (ema-delta)*(2/33)
		}

		// now we can calculate the final average and then deduce the variance
		average = average / float64(len(sets1))
		for _, delta := range deltas {
			variance = variance + (average-delta)*(average-delta)
		}
		variance = variance / float64(len(sets1))

		// add the outcome to the ones tested
		out := model.Outcome{
			Set:      set1,
			Average:  ema,
			Variance: variance,
		}
		outcomes = append(outcomes, out)
	}

	// before picking the top five outcomes, we want to put more weight to
	// outcomes that are closer to other positive outcomes; as we don't care
	// about the absolute value of the score, we simply multiply the average
	// score of the two sets by the similary between them
	weights := make(map[model.Set]float64)
	for _, out1 := range outcomes {
		var weight float64
		for _, out2 := range outcomes {
			distance := model.Distance(out1.Set, out2.Set)
			similarity := 5 - float64(distance)
			weight = weight + similarity*(out1.Average+out2.Average)/2
		}
		weights[out1.Set] = weight
	}
	for i, out := range outcomes {
		outcomes[i].Weight = weights[out.Set]
	}

	// as planned, we now sort all outcomes by their weight in order to
	// determine the top five outcomes of this training; we then use the average
	// variance of these five outcomes as the basis for the score of the
	// strategy; we use the inverse because we want to choose strategies with
	// low variance - the absolute scale doesn't matter here
	var variance float64
	outcomes.Sort(outcome.Reverse(outcome.AverageAsc))
	pick := make([]model.Set, 0, 5)
	for _, out := range outcomes[0:5] {
		variance = variance + out.Variance
		pick = append(pick, out.Set)
	}
	strat := model.Strategy{
		Score: 1 / variance,
		Sets:  pick,
	}

	return strat
}

func (s *Standard) getSignature(wiz1 uint64, wiz2 uint64, nonce1 uint32, nonce2 uint32, commit model.Hash) ([]byte, error) {

	// swap wizards and nonces if first ID is smaller
	if wiz1 > wiz2 {
		wiz1, wiz2 = wiz2, wiz1
		nonce1, nonce2 = nonce2, nonce1
	}

	address := common.HexToAddress(s.tour.Address())

	// create the data to be hashed for the signature
	data := make([]byte, 126)
	data[0] = 0x19               // the hardcoded EIP191 prefix
	data[1] = 0x00               // the hardcoded EIP191 version data
	copy(data[2:22], address[:]) // the contract address
	binary.BigEndian.PutUint64(data[46:54], wiz1)
	binary.BigEndian.PutUint64(data[78:86], wiz2)
	binary.BigEndian.PutUint32(data[86:90], nonce1)
	binary.BigEndian.PutUint32(data[90:94], nonce2)
	copy(data[94:126], commit[:])

	// create the external hash, as used by the contract
	hash := crypto.Keccak256(data)

	// create the internal hash, as used by ethereum EIP 191
	hash = accounts.TextHash(hash)

	// use the private key of our keystore directly to sign
	sig, _ := crypto.Sign(hash, s.key.PrivateKey)

	// no idea why? - OK they do this at the API level at the very end
	sig[64] += 27

	return sig, nil
}

func resolveDuel(set1 model.Set, set2 model.Set, aff1 uint8, aff2 uint8, pow1 uint64, pow2 uint64) int64 {

	// get the duel score
	score := duelScore(set1, set2, int64(aff1), int64(aff2))

	// if the score is positive, get normal power output
	if score > 0 {
		s := big.NewInt(score)
		p1 := new(big.Int).SetUint64(pow1)
		p2 := new(big.Int).SetUint64(pow2)
		return powerTransfer(s, p1, p2)
	}

	// if the score is negative, reverse the input & output
	if score < 0 {
		s := big.NewInt(-score)
		p1 := new(big.Int).SetUint64(pow1)
		p2 := new(big.Int).SetUint64(pow2)
		return -powerTransfer(s, p2, p1)
	}

	return 0
}

func duelScore(set1 model.Set, set2 model.Set, aff1 int64, aff2 int64) int64 {

	// the weight of each round
	weights := []int64{78, 79, 81, 86, 100}

	// for the 5 moves, deduce two each to normalize to zero
	set1[0] -= 2
	set1[1] -= 2
	set1[2] -= 2
	set1[3] -= 2
	set1[4] -= 2
	set2[0] -= 2
	set2[1] -= 2
	set2[2] -= 2
	set2[3] -= 2
	set2[4] -= 2

	// normalize affinities to zero as well
	aff1 -= 2
	aff2 -= 2

	// go through the weights for each round
	var score int64
	for i := 0; i < 5; i++ {

		// get the moves with given index
		move1 := int64(set1[i])
		move2 := int64(set2[i])

		// draw does nothing
		if move1 == move2 {
			continue
		}

		// we want -1 for a loss and +1 for a win of first player
		var round int64
		if (move1 == 0 && move2 == 2) ||
			(move1 == 1 && move2 == 0) ||
			(move1 == 2 && move2 == 1) {
			round = 1
		} else {
			round = -1
		}

		// increase precision
		round = round * 100

		// increase effect with affinities
		if move1 == aff1 {
			round = round * 130 / 100
		}
		if move2 == aff2 {
			round = round * 130 / 100
		}

		score = score + round*weights[i]
	}

	// reduce precision
	score = score / 100

	return score
}

func powerTransfer(s *big.Int, p1 *big.Int, p2 *big.Int) int64 {

	// sum of weights
	sum := big.NewInt(78 + 79 + 81 + 86 + 100)

	// normalize the score
	n1k := big.NewInt(1024)
	math.U256(s.Mul(s, n1k))
	math.U256(s.Div(s, sum))

	// put upper limit
	if s.Cmp(n1k) > 0 {
		s.Set(n1k)
	}

	// stopped following here
	sum.SetInt64(512)
	s = fakePowQ10(s, sum)

	// put upper limit on what we can win
	c1 := big.NewInt(7)
	math.U256(c1.Mul(p1, c1))
	if p2.Cmp(c1) > 0 {
		p2.Set(c1)
	}
	c2 := big.NewInt(7)
	math.U256(c2.Mul(p2, c2))
	if p1.Cmp(c2) > 0 {
		p1.Set(c2)
	}

	math.U256(n1k.Mul(n1k, p2))
	math.U256(n1k.Div(n1k, p1))
	s = fakePowQ10(s, n1k)

	math.U256(s.Mul(s, p2))
	math.U256(s.Rsh(s, 10))
	if !s.IsInt64() {
		panic("overflow for score result")
	}

	return s.Int64()
}

func fakePowQ10(x *big.Int, y *big.Int) *big.Int {
	tmp := big.NewInt(8)
	math.U256(y.Add(y, tmp))
	math.U256(y.Rsh(y, 4))
	tmp.SetInt64(64)
	return fakePowInternal(x, y, tmp, 5)
}

func fakePowInternal(x *big.Int, y *big.Int, den *big.Int, its uint) *big.Int {

	tmp := new(big.Int)

	res := big.NewInt(1024)
	exp := math.U256(new(big.Int).Div(y, den))

	num := big.NewInt(22)
	for exp.Cmp(num) >= 0 {
		math.U256(tmp.Exp(x, num, nil))
		math.U256(res.Mul(res, tmp))
		math.U256(res.Rsh(res, 220))
		math.U256(exp.Sub(exp, num))
	}

	num.SetInt64(0)
	if exp.Cmp(num) > 0 {
		math.U256(num.Exp(x, exp, nil))
		math.U256(res.Mul(res, num))
		num.SetInt64(10)
		math.U256(num.Mul(exp, num))
		math.U256(res.Rsh(res, uint(num.Uint64())))
	}

	num.SetInt64(0)
	fra := math.U256(new(big.Int).Mod(y, den))
	if its == 0 || fra.Cmp(num) == 0 {
		return res
	}

	num.SetInt64(1024)
	math.U256(x.Sub(num, x))
	tmp = fakePowInternal(x, den, fra, its-1)
	math.U256(tmp.Sub(num, tmp))
	math.U256(res.Mul(res, tmp))
	math.U256(res.Rsh(res, 10))

	return res
}

func generateSecret(chaID string, set model.Set, wizID uint64) model.Secret {

	// geneerate a random salt
	var salt model.Hash
	_, _ = rand.Read(salt[:]) // never returns error

	// generate the commit
	data := make([]byte, 64)
	copy(data[0:5], set[:])
	copy(data[32:64], salt[:])
	commit := crypto.Keccak256Hash(data)

	// build the secret
	secret := model.Secret{
		ChaID:  chaID,
		WizID:  wizID,
		Set:    set,
		Salt:   salt,
		Commit: model.Hash(commit),
	}

	return secret
}
