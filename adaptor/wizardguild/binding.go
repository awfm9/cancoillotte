// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package wizardguild

import (
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/awishformore/cancoillotte/bindings"
)

type Binding struct {
	call     *bind.CallOpts
	contract *bindings.WizardGuild
}

func NewBinding(address string, cli *ethclient.Client) (*Binding, error) {
	contract, _ := bindings.NewWizardGuild(common.HexToAddress(address), cli)
	b := &Binding{
		call:     &bind.CallOpts{},
		contract: contract,
	}
	_, err := b.Owner(6139)
	if err != nil {
		return nil, errors.Wrap(err, "could not check status")
	}
	return b, nil
}

func (b *Binding) Contract() *bindings.WizardGuild {
	return b.contract
}

func (b *Binding) Owner(wizID uint64) (string, error) {
	owner, err := b.contract.OwnerOf(b.call,
		big.NewInt(0).SetUint64(wizID),
	)
	if err != nil {
		return "", errors.Wrap(err, "could not call contract")
	}
	return owner.Hex(), nil
}
