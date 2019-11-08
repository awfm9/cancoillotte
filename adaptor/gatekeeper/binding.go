// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package gatekeeper

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/awishformore/cancoillotte/bindings"
)

type Binding struct {
	call     *bind.CallOpts
	auth     *bind.TransactOpts
	contract *bindings.InauguralGateKeeper
}

func NewBinding(address string, cli *ethclient.Client, auth *bind.TransactOpts) (*Binding, error) {
	contract, _ := bindings.NewInauguralGateKeeper(common.HexToAddress(address), cli)
	nonce, err := cli.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, errors.Wrap(err, "could not get pending nonce")
	}
	auth.Nonce = big.NewInt(0).SetUint64(nonce)
	b := &Binding{
		call:     &bind.CallOpts{},
		auth:     auth,
		contract: contract,
	}
	_, _, err = b.Costs()
	if err != nil {
		return nil, errors.Wrap(err, "could not check status")
	}
	return b, nil
}

func (b *Binding) Contract() *bindings.InauguralGateKeeper {
	return b.contract
}

func (b *Binding) Costs() (uint64, uint64, error) {
	costs, err := b.contract.WizardCosts(b.call)
	if err != nil {
		return 0, 0, errors.Wrap(err, "could not call contract")
	}
	neutral := costs.NeutralWizardCost.Uint64()
	elemental := costs.ElementalWizardCost.Uint64()
	return neutral, elemental, nil
}

func (b *Binding) Conjure(aff uint8, cost uint64) (string, error) {
	hash, err := b.exec(b.conjure(aff), cost)
	if err != nil {
		return "", errors.Wrap(err, "could not execute conjure")
	}
	return hash, nil
}

func (b *Binding) conjure(aff uint8) func() (string, error) {
	return func() (string, error) {
		tx, err := b.contract.ConjureWizard(b.auth, aff)
		if err != nil {
			return "", errors.Wrap(err, "could not conjure wizard")
		}
		return tx.Hash().Hex(), nil
	}
}

func (b *Binding) exec(call func() (string, error), cost uint64) (string, error) {
	defer func() {
		b.auth.Value = nil
		b.auth.Nonce.Add(b.auth.Nonce, big.NewInt(1))
	}()
	b.auth.Value = big.NewInt(0).SetUint64(cost)
	hash, err := call()
	if err != nil {
		return "", errors.Wrap(err, "could not execute call")
	}
	return hash, nil
}
