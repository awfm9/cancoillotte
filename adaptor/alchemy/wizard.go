// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package alchemy

import (
	"strconv"

	"github.com/awishformore/cancoillotte/model"
)

type WizardResponse struct {
	Wizards []Wizard `json:"wizards"`
}

type Wizard struct {
	ID                    string  `json:"id"`
	Owner                 string  `json:"owner"`
	Affinity              uint8   `json:"affinity"`
	InitialPower          string  `json:"initialPower"`
	Power                 string  `json:"power"`
	EliminatedBlockNumber *uint64 `json:"eliminatedBlockNumber"`
	CreatedBlockNumber    uint64  `json:"createdBlockNumber"`
}

func (w Wizard) Convert() model.Wizard {
	id, _ := strconv.ParseUint(w.ID, 10, 64)
	pow, _ := strconv.ParseUint(w.Power, 10, 64)
	wiz := model.Wizard{
		ID:    id,
		Owner: w.Owner,
		Aff:   w.Affinity,
		Pow:   pow,
	}
	return wiz
}
