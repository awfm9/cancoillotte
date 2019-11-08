// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package tournament

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/awishformore/cancoillotte/bindings"
	"github.com/awishformore/cancoillotte/model"
)

type Binding struct {
	address  string
	call     *bind.CallOpts
	auth     *bind.TransactOpts
	cli      *ethclient.Client
	contract *bindings.BasicTournament
}

func NewBinding(address string, cli *ethclient.Client, auth *bind.TransactOpts) (*Binding, error) {
	contract, _ := bindings.NewBasicTournament(common.HexToAddress(address), cli)
	b := &Binding{
		address:  address,
		call:     &bind.CallOpts{},
		auth:     auth,
		cli:      cli,
		contract: contract,
	}
	_, err := b.Remaining()
	if err != nil {
		return nil, errors.Wrap(err, "could not check status")
	}
	return b, nil
}

func (b *Binding) Address() string {
	return b.address
}

func (b *Binding) Contract() *bindings.BasicTournament {
	return b.contract
}

func (b *Binding) Remaining() (uint, error) {
	remaining, err := b.contract.GetRemainingWizards(b.call)
	if err != nil {
		return 0, errors.Wrap(err, "could not call contract")
	}
	return uint(remaining.Uint64()), nil
}

func (b *Binding) Ascending() (uint64, error) {
	ascending, err := b.contract.GetAscendingWizardId(b.call)
	if err != nil {
		return 0, errors.Wrap(err, "could not call contract")
	}
	return ascending.Uint64(), nil
}

func (b *Binding) Ready(wizID uint64) (bool, error) {
	ready, err := b.contract.IsReady(b.call,
		big.NewInt(0).SetUint64(wizID),
	)
	if err != nil {
		return false, errors.Wrap(err, "could not call contract")
	}
	return ready, nil
}

func (b *Binding) Wizard(wizID uint64) (model.Wizard, error) {
	wizard, err := b.contract.GetWizard(b.call,
		big.NewInt(0).SetUint64(wizID),
	)
	if err != nil {
		return model.Wizard{}, errors.Wrap(err, "could not call contract")
	}
	w := model.Wizard{
		ID:    wizID,
		Owner: "",
		Aff:   uint8(wizard.Affinity.Uint64()),
		Pow:   wizard.Power.Uint64(),
		Nonce: uint32(wizard.Nonce.Uint64()),
	}
	return w, nil
}

func (b *Binding) Paused() (bool, error) {
	paused, err := b.contract.IsPaused(b.call)
	if err != nil {
		return false, errors.Wrap(err, "could not call contract")
	}
	return paused, nil
}

func (b *Binding) Height() (uint64, error) {
	header, err := b.cli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, errors.Wrap(err, "could not get header")
	}
	return header.Number.Uint64(), nil
}

func (b *Binding) Mold() (Mold, error) {
	moldStart, sessionDuration, doublingDuration, basePower, err := b.contract.GetBlueMoldParameters(b.call)
	if err != nil {
		return Mold{}, errors.Wrap(err, "could not call contract")
	}
	params := Mold{
		BasePower:        basePower.Uint64(),
		MoldStart:        moldStart.Uint64(),
		SessionDuration:  sessionDuration.Uint64(),
		DoublingDuration: doublingDuration.Uint64(),
	}
	return params, nil
}

func (b *Binding) Time() (Time, error) {
	params, err := b.contract.GetTimeParameters(b.call)
	if err != nil {
		return Time{}, errors.Wrap(err, "could not call contract")
	}
	time := Time{
		Start:     params.TournamentStartBlock.Uint64(),
		Pause:     params.PauseEndedBlock.Uint64(),
		Admission: params.AdmissionDuration.Uint64(),
		Revival:   params.RevivalDuration.Uint64(),
		AscStart:  params.AscensionWindowStart.Uint64(),
		AscDur:    params.AscensionWindowDuration.Uint64(),
		FigStart:  params.FightWindowStart.Uint64(),
		FigDur:    params.FightWindowDuration.Uint64(),
		ResStart:  params.ResolutionWindowStart.Uint64(),
		ResDur:    params.ResolutionWindowDuration.Uint64(),
		CulStart:  params.CullingWindowStart.Uint64(),
		CulDur:    params.CullingWindowDuration.Uint64(),
	}
	return time, nil
}

func (b *Binding) Timeout(wiz1 uint64, wiz2 uint64, price uint64) (string, error) {
	hash, err := b.exec(b.timeout(wiz1, wiz2), price)
	if err != nil {
		return "", errors.Wrap(err, "could not execute timeout")
	}
	return hash, nil
}

func (b *Binding) timeout(wiz1 uint64, wiz2 uint64) func() (string, error) {
	return func() (string, error) {
		tx, err := b.contract.ResolveTimedOutDuel(b.auth,
			big.NewInt(0).SetUint64(wiz1),
			big.NewInt(0).SetUint64(wiz2),
		)
		if err != nil {
			return "", errors.Wrap(err, "could not timeout duel")
		}
		return tx.Hash().Hex(), nil
	}
}

func (b *Binding) exec(call func() (string, error), price uint64) (string, error) {
	defer func() {
		b.auth.GasPrice = nil
		b.auth.Nonce.Add(b.auth.Nonce, big.NewInt(1))
	}()
	b.auth.GasPrice = big.NewInt(0).SetUint64(price)
	hash, err := call()
	if err != nil {
		return "", errors.Wrap(err, "could not execute call")
	}
	return hash, nil
}
