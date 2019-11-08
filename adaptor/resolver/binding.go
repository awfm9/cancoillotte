// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package resolver

import (
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/awishformore/cancoillotte/bindings"
	"github.com/awishformore/cancoillotte/model"
)

type Binding struct {
	call     *bind.CallOpts
	contract *bindings.ThreeAffinityDuelResolver
}

func NewBinding(address string, cli *ethclient.Client) (*Binding, error) {
	contract, _ := bindings.NewThreeAffinityDuelResolver(common.HexToAddress(address), cli)
	b := &Binding{
		call:     &bind.CallOpts{},
		contract: contract,
	}
	_, err := b.Valid(model.Set{})
	if err != nil {
		return nil, errors.Wrap(err, "could not check status")
	}
	return b, nil
}

func (b *Binding) Contract() *bindings.ThreeAffinityDuelResolver {
	return b.contract
}

func (b *Binding) Valid(set model.Set) (bool, error) {
	valid, err := b.contract.IsValidMoveSet(b.call, set.Full())
	if err != nil {
		return false, errors.Wrap(err, "could not call contract")
	}
	return valid, nil
}

func (b *Binding) Score(set1 model.Set, set2 model.Set, aff1 uint8, aff2 uint8, pow1 uint64, pow2 uint64) (int64, error) {
	delta, err := b.contract.ResolveDuel(b.call,
		set1.Full(),
		set2.Full(),
		big.NewInt(0).SetUint64(pow1),
		big.NewInt(0).SetUint64(pow2),
		big.NewInt(int64(aff1)),
		big.NewInt(int64(aff2)),
	)
	if err != nil {
		return 0, errors.Wrap(err, "could not call contract")
	}
	return delta.Int64(), nil
}
