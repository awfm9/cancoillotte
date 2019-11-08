// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package core

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/awishformore/cancoillotte/bindings"
	"github.com/awishformore/cancoillotte/model"
)

type Subscriber struct {
	log   zerolog.Logger
	cli   *ethclient.Client
	guild *bindings.WizardGuild
	tour  *bindings.BasicTournament
	strat Strategy
	wg    *sync.WaitGroup
	subs  []ethereum.Subscription
}

func NewSubscriber(log zerolog.Logger, cli *ethclient.Client, guild *bindings.WizardGuild, tour *bindings.BasicTournament, strat Strategy) (*Subscriber, error) {
	s := &Subscriber{
		cli:   cli,
		log:   log,
		guild: guild,
		tour:  tour,
		strat: strat,
		wg:    &sync.WaitGroup{},
	}

	subscribes := [](func() (ethereum.Subscription, error)){
		s.Blocks,
		s.Creations,
		s.Transfers,
		s.Challenges,
		s.Withdrawals,
		s.Fights,
		s.Reveals,
		s.Results,
		s.Timeouts,
		s.Eliminations,
		s.Attempts,
		s.Pairs,
		s.Contests,
		s.Successes,
		s.Revivals,
		s.Wipes,
		s.Claims,
	}
	for _, subscribe := range subscribes {
		sub, err := subscribe()
		if err != nil {
			return nil, errors.Wrap(err, "could not subscribe")
		}
		s.subs = append(s.subs, sub)
	}

	s.wg.Add(len(s.subs))

	return s, nil
}

// Blocks will subscribe to new header events.
func (s *Subscriber) Blocks() (ethereum.Subscription, error) {

	blocks := make(chan *types.Header, 1)
	sub, err := s.cli.SubscribeNewHead(context.Background(), blocks)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to header events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case block := <-blocks:

				hash := block.Hash().Hex()
				height := block.Number.Uint64()

				log := s.log.With().
					Str("event", "header").
					Str("hash", hash).
					Uint64("height", height).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnBlock(hash, height)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Creations will subscribe to new creation events.
func (s *Subscriber) Creations() (ethereum.Subscription, error) {

	creations := make(chan *bindings.WizardGuildWizardConjured, 1)
	sub, err := s.guild.WatchWizardConjured(&bind.WatchOpts{}, creations)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to creation events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case creation := <-creations:

				wiz := creation.WizardId.Uint64()
				aff := creation.Affinity
				pow := creation.InnatePower.Uint64()

				log := s.log.With().
					Str("event", "creation").
					Uint64("wiz", wiz).
					Uint8("aff", aff).
					Uint64("pow", pow).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnCreation(wiz, aff, pow)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Transfers will subscribe to new transfer events.
func (s *Subscriber) Transfers() (ethereum.Subscription, error) {

	transfers := make(chan *bindings.WizardGuildTransfer, 1)
	sub, err := s.guild.WatchTransfer(&bind.WatchOpts{}, transfers, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to transfer events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case transfer := <-transfers:

				wiz := transfer.TokenId.Uint64()
				from := transfer.From.Hex()
				to := transfer.To.Hex()

				log := s.log.With().
					Str("event", "transfer").
					Uint64("wiz", wiz).
					Str("from", from).
					Str("to", to).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnTransfer(wiz, from, to)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Challenges will subscribe to new challenge events.
func (s *Subscriber) Challenges() (ethereum.Subscription, error) {

	challenges := make(chan *bindings.BasicTournamentOneSidedCommitAdded, 1)
	sub, err := s.tour.WatchOneSidedCommitAdded(&bind.WatchOpts{}, challenges)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to challenge events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case challenge := <-challenges:

				wiz1 := challenge.CommittingWizardId.Uint64()
				wiz2 := challenge.OtherWizardId.Uint64()
				nonce1 := challenge.CommittingWizardNonce.Uint64()
				nonce2 := challenge.OtherWizardNonce.Uint64()
				commit1 := challenge.Commitment

				log := s.log.With().
					Str("event", "challenge").
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Uint64("nonce1", nonce1).
					Uint64("nonce2", nonce2).
					Hex("commit1", commit1[:]).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnChallenge(wiz1, wiz2, nonce1, nonce2, commit1)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Withdrawals will subscribe to withdrawals of challenges.
func (s *Subscriber) Withdrawals() (ethereum.Subscription, error) {

	withdrawals := make(chan *bindings.BasicTournamentOneSidedCommitCancelled, 1)
	sub, err := s.tour.WatchOneSidedCommitCancelled(&bind.WatchOpts{}, withdrawals)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to withdrawal events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case withdrawal := <-withdrawals:

				wiz := withdrawal.WizardId.Uint64()

				log := s.log.With().
					Str("event", "withdrawal").
					Uint64("wiz", wiz).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnWithdrawal(wiz)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Fights will subscribe to new fight events.
func (s *Subscriber) Fights() (ethereum.Subscription, error) {

	fights := make(chan *bindings.BasicTournamentDuelStart, 1)
	sub, err := s.tour.WatchDuelStart(&bind.WatchOpts{}, fights)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to fight events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case fight := <-fights:

				duel := fight.DuelId
				wiz1 := fight.WizardId1.Uint64()
				wiz2 := fight.WizardId2.Uint64()
				timeout := fight.TimeoutBlock.Uint64()
				contest := fight.IsAscensionBattle
				commit1 := fight.Commit1
				commit2 := fight.Commit2

				log := s.log.With().
					Str("event", "fight").
					Hex("duel", duel[:]).
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Uint64("timeout", timeout).
					Bool("contest", contest).
					Hex("commit1", commit1[:]).
					Hex("commit2", commit2[:]).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnFight(duel, wiz1, wiz2, timeout, contest, commit1, commit2)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Reveals will subscribe to new reveal events.
func (s *Subscriber) Reveals() (ethereum.Subscription, error) {

	reveals := make(chan *bindings.BasicTournamentOneSidedRevealAdded, 1)
	sub, err := s.tour.WatchOneSidedRevealAdded(&bind.WatchOpts{}, reveals)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to reveal events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case reveal := <-reveals:

				duel := reveal.DuelId
				wiz1 := reveal.CommittingWizardId.Uint64()
				wiz2 := reveal.OtherWizardId.Uint64()

				log := s.log.With().
					Str("event", "reveal").
					Hex("duel", duel[:]).
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnReveal(duel, wiz1, wiz2)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Results will subscribe to new result events.
func (s *Subscriber) Results() (ethereum.Subscription, error) {

	results := make(chan *bindings.BasicTournamentDuelEnd, 1)
	sub, err := s.tour.WatchDuelEnd(&bind.WatchOpts{}, results)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to result events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case result := <-results:

				duel := result.DuelId
				wiz1 := result.WizardId1.Uint64()
				wiz2 := result.WizardId2.Uint64()
				set1 := model.Compact(result.MoveSet1)
				set2 := model.Compact(result.MoveSet2)
				pow1 := result.Power1.Uint64()
				pow2 := result.Power2.Uint64()

				log := s.log.With().
					Str("event", "result").
					Hex("duel", duel[:]).
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Hex("set1", set1[:]).
					Hex("set2", set2[:]).
					Uint64("pow1", pow1).
					Uint64("pow2", pow2).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnResult(duel, wiz1, wiz2, set1, set2, pow1, pow2)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Timeouts will subscribe to new timeout events.
func (s *Subscriber) Timeouts() (ethereum.Subscription, error) {

	timeouts := make(chan *bindings.BasicTournamentDuelTimeOut, 1)
	sub, err := s.tour.WatchDuelTimeOut(&bind.WatchOpts{}, timeouts)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to timeout events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case timeout := <-timeouts:

				duel := timeout.DuelId
				wiz1 := timeout.WizardId1.Uint64()
				wiz2 := timeout.WizardId2.Uint64()
				pow1 := timeout.Power1.Uint64()
				pow2 := timeout.Power2.Uint64()

				log := s.log.With().
					Str("event", "timeout").
					Hex("duel", duel[:]).
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Uint64("pow1", pow1).
					Uint64("pow2", pow2).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnTimeout(duel, wiz1, wiz2, pow1, pow2)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Eliminations will subscribe to new elimination events.
func (s *Subscriber) Eliminations() (ethereum.Subscription, error) {

	elims := make(chan *bindings.BasicTournamentWizardElimination, 1)
	sub, err := s.tour.WatchWizardElimination(&bind.WatchOpts{}, elims)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to elimination events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case elim := <-elims:

				wiz := elim.WizardId.Uint64()

				log := s.log.With().
					Str("event", "elimination").
					Uint64("wiz", wiz).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnElimination(wiz)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Attempts will subscribe to new attempt events.
func (s *Subscriber) Attempts() (ethereum.Subscription, error) {

	attempts := make(chan *bindings.BasicTournamentAscensionStart, 1)
	sub, err := s.tour.WatchAscensionStart(&bind.WatchOpts{}, attempts)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to attempt events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case attempt := <-attempts:

				wiz := attempt.WizardId.Uint64()

				log := s.log.With().
					Str("event", "attempt").
					Uint64("wiz", wiz).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnAttempt(wiz)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Pairs will subscribe to new pair events.
func (s *Subscriber) Pairs() (ethereum.Subscription, error) {

	pairs := make(chan *bindings.BasicTournamentAscensionPairUp, 1)
	sub, err := s.tour.WatchAscensionPairUp(&bind.WatchOpts{}, pairs)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to attempt events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case pair := <-pairs:

				wiz1 := pair.WizardId1.Uint64()
				wiz2 := pair.WizardId2.Uint64()

				log := s.log.With().
					Str("event", "pair").
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnPair(wiz1, wiz2)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Contests will subscribe to new contest events.
func (s *Subscriber) Contests() (ethereum.Subscription, error) {

	contests := make(chan *bindings.BasicTournamentAscensionChallenged, 1)
	sub, err := s.tour.WatchAscensionChallenged(&bind.WatchOpts{}, contests)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to contest events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case contest := <-contests:

				wiz1 := contest.AscendingWizardId.Uint64()
				wiz2 := contest.ChallengingWizardId.Uint64()
				commit2 := contest.Commitment

				log := s.log.With().
					Str("event", "contest").
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Hex("commit2", commit2[:]).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnContest(wiz1, wiz2, commit2)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Successes will subscribe to new success events.
func (s *Subscriber) Successes() (ethereum.Subscription, error) {

	successes := make(chan *bindings.BasicTournamentAscensionComplete, 1)
	sub, err := s.tour.WatchAscensionComplete(&bind.WatchOpts{}, successes)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to success events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case success := <-successes:

				wiz := success.WizardId.Uint64()
				pow := success.Power.Uint64()

				log := s.log.With().
					Str("event", "success").
					Uint64("wiz", wiz).
					Uint64("pow", pow).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnSuccess(wiz, pow)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Revivals will subscribe to new revival events.
func (s *Subscriber) Revivals() (ethereum.Subscription, error) {

	revivals := make(chan *bindings.BasicTournamentRevive, 1)
	sub, err := s.tour.WatchRevive(&bind.WatchOpts{}, revivals)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to revival events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case revival := <-revivals:

				wiz := revival.WizId.Uint64()
				pow := revival.Power.Uint64()

				log := s.log.With().
					Str("event", "revival").
					Uint64("wiz", wiz).
					Uint64("pow", pow).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnRevival(wiz, pow)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Wipes will subscribe to new wipe events.
func (s *Subscriber) Wipes() (ethereum.Subscription, error) {

	wipes := make(chan *bindings.BasicTournamentPowerTransferred, 1)
	sub, err := s.tour.WatchPowerTransferred(&bind.WatchOpts{}, wipes)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to wipe events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case wipe := <-wipes:

				wiz1 := wipe.SendingWizId.Uint64()
				wiz2 := wipe.ReceivingWizId.Uint64()
				pow1 := wipe.AmountTransferred.Uint64()
				reason := wipe.Reason

				log := s.log.With().
					Str("event", "wipe").
					Uint64("wiz1", wiz1).
					Uint64("wiz2", wiz2).
					Uint64("pow1", pow1).
					Uint8("reason", reason).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnWipe(wiz1, wiz2, pow1, reason)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Claims will subscribe to new claim events.
func (s *Subscriber) Claims() (ethereum.Subscription, error) {

	claims := make(chan *bindings.BasicTournamentPrizeClaimed, 1)
	sub, err := s.tour.WatchPrizeClaimed(&bind.WatchOpts{}, claims)
	if err != nil {
		return nil, errors.Wrap(err, "could not subscribe to claim events")
	}

	go func() {
		defer s.wg.Done()
	Loop:
		for {
			select {
			case claim := <-claims:

				wiz := claim.ClaimingWinnerId.Uint64()
				prize := claim.PrizeAmount.Uint64()

				log := s.log.With().
					Str("event", "claim").
					Uint64("wiz", wiz).
					Uint64("prize", prize).
					Logger()

				log.Debug().Msg("event received")

				err := s.strat.OnClaim(wiz, prize)
				if err != nil {
					log.Error().Err(err).Msg("processing failed")
					continue
				}

			case <-sub.Err():
				break Loop
			}
		}
	}()

	return sub, nil
}

// Stop will stop the subscriber.
func (s *Subscriber) Stop() {
	for _, sub := range s.subs {
		sub.Unsubscribe()
	}
	s.wg.Wait()
}
