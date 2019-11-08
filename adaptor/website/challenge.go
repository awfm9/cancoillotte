// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

import (
	"strconv"
	"time"

	"github.com/awishformore/cancoillotte/model"
	"github.com/ethereum/go-ethereum/common"
)

type Challenge struct {
	ID           string    `json:"id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	LastActionAt time.Time `json:"lastActionAt"`
	HomeWizard   Wizard    `json:"homeWizard"`
	AwayWizard   Wizard    `json:"awayWizard"`
	Outcome      Outcome   `json:"outcome"`
}

func (c Challenge) Convert() model.Challenge {
	wiz1, _ := strconv.ParseUint(c.HomeWizard.ID, 10, 64)
	wiz2, _ := strconv.ParseUint(c.AwayWizard.ID, 10, 64)
	challenge := model.Challenge{
		ID:      c.ID,
		Wiz1:    wiz1,
		Wiz2:    wiz2,
		Owner1:  common.HexToAddress(c.HomeWizard.OwnerAddress).Hex(),
		Owner2:  common.HexToAddress(c.AwayWizard.OwnerAddress).Hex(),
		Player1: c.HomeWizard.Owner.ID,
		Player2: c.AwayWizard.Owner.ID,
		Online1: c.HomeWizard.Owner.Status == "ONLINE",
		Online2: c.AwayWizard.Owner.Status == "ONLINE",
		Status:  c.Status,
	}
	return challenge
}
