// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package duels

import (
	"encoding/hex"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/awishformore/cancoillotte/adaptor/alchemy"
	"github.com/awishformore/cancoillotte/adaptor/database"
	"github.com/awishformore/cancoillotte/model"
	"github.com/awishformore/cancoillotte/state/wizards"
)

type Repo struct {
	sync.Mutex
	alch  *alchemy.Client
	kv    *database.KV
	duels map[string]model.Duel
}

func NewRepo(log zerolog.Logger, alch *alchemy.Client, kv *database.KV, wiz *wizards.Repo) (*Repo, error) {
	duels := make(map[string]model.Duel)
	alDuels, err := alch.Duels()
	if err != nil {
		return nil, errors.Wrap(err, "could not bootstrap duels")
	}
	for _, duel := range alDuels {
		wiz1 := wiz.Get(duel.Wiz1)
		wiz2 := wiz.Get(duel.Wiz2)
		duel.Owner1 = wiz1.Owner
		duel.Owner2 = wiz2.Owner
		duel.Player1 = wiz1.Player
		duel.Player2 = wiz2.Player
		err = kv.Duel(duel)
		if err != nil {
			return nil, errors.Wrap(err, "could not save duel to disk")
		}
		key := hex.EncodeToString(duel.ID[:])
		duels[key] = duel
	}
	log.Info().Int("duels", len(duels)).Msg("duels loaded from alchemy")
	r := &Repo{
		alch:  alch,
		kv:    kv,
		duels: duels,
	}
	return r, nil
}

func (r *Repo) Update(duelID model.Hash, update func(duel *model.Duel)) error {
	r.Lock()
	defer r.Unlock()
	key := hex.EncodeToString(duelID[:])
	duel, ok := r.duels[key]
	if !ok {
		duel = model.Duel{}
	}
	original := duel
	update(&duel)
	if (duel == original) || (duel.ID == model.Hash{}) {
		return nil
	}
	err := r.kv.Duel(duel)
	if err != nil {
		return errors.Wrap(err, "could not save duel to disk")
	}
	r.duels[key] = duel
	return nil
}

func (r *Repo) Select() model.DuelList {
	r.Lock()
	defer r.Unlock()
	duels := make(model.DuelList, 0, len(r.duels))
	for _, duel := range r.duels {
		duels = append(duels, duel)
	}
	return duels
}
