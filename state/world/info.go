// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package world

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/awishformore/cancoillotte/adaptor/tournament"
	"github.com/awishformore/cancoillotte/state/duels"
	"github.com/awishformore/cancoillotte/state/players"
)

type Info struct {
	sync.Mutex
	cli     *ethclient.Client
	tour    *tournament.Binding
	online  map[string]bool
	action  map[string]uint64
	address string
}

func NewInfo(cli *ethclient.Client, tour *tournament.Binding, duel *duels.Repo, play *players.Repo, address common.Address) (*Info, error) {
	action := make(map[string]uint64)
	duels := duel.Select()
	for _, duel := range duels {
		player1, err := play.Get(duel.Owner1, duel.Wiz1)
		if err != nil {
			return nil, errors.Wrap(err, "could not get home player")
		}
		player2, err := play.Get(duel.Owner2, duel.Wiz1)
		if err != nil {
			return nil, errors.Wrap(err, "could not get away player")
		}
		if duel.End > action[player1] {
			action[player1] = duel.End
		}
		if duel.End > action[player2] {
			action[player2] = duel.End
		}
	}
	i := &Info{
		cli:     cli,
		tour:    tour,
		online:  make(map[string]bool),
		action:  action,
		address: address.Hex(),
	}
	return i, nil
}

func (i *Info) Address() string {
	return i.address
}

func (i *Info) Balance() (uint64, error) {
	balance, err := i.cli.PendingBalanceAt(context.Background(), common.HexToAddress(i.address))
	if err != nil {
		return 0, errors.Wrap(err, "could not retrieve balance")
	}
	return balance.Uint64(), nil
}

func (i *Info) Reset() {
	i.Lock()
	defer i.Unlock()
	i.online = make(map[string]bool)
}

func (i *Info) Online(player string) {
	i.Lock()
	defer i.Unlock()
	i.online[player] = true
}

func (i *Info) Action(player string, height uint64) {
	i.Lock()
	defer i.Unlock()
	i.action[player] = height
}

func (i *Info) Active(player string, height uint64) bool {
	i.Lock()
	defer i.Unlock()
	// consider players active if they played within the last session
	return i.action[player] > height-2150
}

func (i *Info) Phase() (string, error) {
	paused, err := i.tour.Paused()
	if err != nil {
		return "", errors.Wrap(err, "could not check paused")
	}
	if paused {
		return Inactive, nil
	}
	height, err := i.tour.Height()
	if err != nil {
		return "", errors.Wrap(err, "could not get current height")
	}
	params, err := i.tour.Time()
	if err != nil {
		return "", errors.Wrap(err, "could not get time parameters")
	}
	if params.Start > height {
		return Inactive, nil
	}
	if params.Start+params.Admission > height {
		return Admission, nil
	}
	if params.Start+params.Admission+params.Revival > height {
		return Revival, nil
	}
	return Elimination, nil
}

func (i *Info) Window() (string, error) {
	paused, err := i.tour.Paused()
	if err != nil {
		return "", errors.Wrap(err, "could not check paused")
	}
	if paused {
		return Idle, nil
	}
	height, err := i.tour.Height()
	if err != nil {
		return "", errors.Wrap(err, "could not get current height")
	}
	params, err := i.tour.Time()
	if err != nil {
		return "", errors.Wrap(err, "could not get time parameters")
	}
	sessionDur := params.AscDur + params.FigDur + params.ResDur + params.CulDur
	if isInWindow(height, params.AscStart, params.AscDur, sessionDur) {
		return Ascension, nil
	}
	if isInWindow(height, params.FigStart, params.FigDur, sessionDur) {
		return Fight, nil
	}
	if isInWindow(height, params.ResStart, params.ResDur, sessionDur) {
		return Resolution, nil
	}
	if isInWindow(height, params.CulStart, params.CulDur, sessionDur) {
		return Culling, nil
	}
	return Idle, nil
}

func isInWindow(height uint64, winStart uint64, winDur uint64, sessDur uint64) bool {
	if height < winStart {
		return false
	}
	offset := (height - winStart) % sessDur
	return offset < winDur
}
