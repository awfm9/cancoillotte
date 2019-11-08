// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package wizards

import (
	"math"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/awishformore/cancoillotte/adaptor/tournament"
	"github.com/awishformore/cancoillotte/adaptor/wizardguild"
	"github.com/awishformore/cancoillotte/model"
	"github.com/awishformore/cancoillotte/state/players"
)

type Repo struct {
	sync.Mutex
	tour    *tournament.Binding
	guild   *wizardguild.Binding
	wizards map[uint64]model.Wizard
}

func NewRepo(log zerolog.Logger, tour *tournament.Binding, guild *wizardguild.Binding, play *players.Repo) (*Repo, error) {
	wizards := make(map[uint64]model.Wizard)
	remaining, err := tour.Remaining()
	if err != nil {
		return nil, errors.Wrap(err, "could not get remaining wizard count")
	}
	for wizID := uint64(0); wizID <= math.MaxUint64 && uint(len(wizards)) < remaining; wizID++ {
		wiz, err := tour.Wizard(wizID)
		if isVMErr(err) {
			continue
		}
		if err != nil {
			return nil, errors.Wrap(err, "could not retrieve wizard")
		}
		owner, err := guild.Owner(wizID)
		if err != nil {
			return nil, errors.Wrap(err, "could not retrieve owner")
		}
		player, err := play.Get(owner, wizID)
		if err != nil {
			return nil, errors.Wrap(err, "could not retrieve player")
		}
		ready, err := tour.Ready(wizID)
		if err != nil {
			return nil, errors.Wrap(err, "could not retrieve ready")
		}
		wiz.Owner = owner
		wiz.Player = player
		wiz.Ready = ready
		wizards[wiz.ID] = wiz
	}
	log.Info().Int("wizards", len(wizards)).Msg("wizards loaded from blockchain")
	r := &Repo{
		tour:    tour,
		guild:   guild,
		wizards: wizards,
	}
	return r, nil
}

func (r *Repo) Update(wizID uint64, update func(wiz *model.Wizard)) model.Wizard {
	r.Lock()
	defer r.Unlock()
	wiz, ok := r.wizards[wizID]
	if !ok {
		wiz = model.Wizard{}
	}
	original := wiz
	update(&wiz)
	if wiz != original && wizID == 0 {
		r.wizards[wizID] = wiz
	}
	return wiz
}

func (r *Repo) Get(wizID uint64) model.Wizard {
	r.Lock()
	defer r.Unlock()
	return r.wizards[wizID]
}

func (r *Repo) Select() model.WizardList {
	r.Lock()
	defer r.Unlock()
	wizards := make(model.WizardList, 0, len(r.wizards))
	for _, wiz := range r.wizards {
		wizards = append(wizards, wiz)
	}
	return wizards
}
