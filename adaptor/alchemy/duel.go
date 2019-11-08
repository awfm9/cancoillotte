// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package alchemy

import (
	"encoding/hex"
	"strconv"

	"github.com/awishformore/cancoillotte/model"
)

type DuelResponse struct {
	Duels []Duel `json:"duels"`
}

type Duel struct {
	ID                string `json:"id"`
	Wizard1ID         string `json:"wizard1Id"`
	Wizard2ID         string `json:"wizard2Id"`
	Affinity1         uint8  `json:"affinity1"`
	Affinity2         uint8  `json:"affinity2"`
	StartPower1       string `json:"startPower1"`
	StartPower2       string `json:"startPower2"`
	EndPower1         string `json:"endPower1"`
	EndPower2         string `json:"endPower2"`
	MoveSet1          string `json:"moveSet1"`
	MoveSet2          string `json:"moveSet2"`
	StartBlock        uint64 `json:"startBlock"`
	EndBlock          uint64 `json:"endBlock"`
	TimeoutBlock      uint64 `json:"timeoutBlock"`
	TimedOut          bool   `json:"timedOut"`
	IsAscensionBattle bool   `json:"isAscensionBattle"`
}

func (d Duel) Convert() model.Duel {
	wiz1, _ := strconv.ParseUint(d.Wizard1ID, 10, 64)
	wiz2, _ := strconv.ParseUint(d.Wizard2ID, 10, 64)
	var id model.Hash
	var set1 model.Set
	var set2 model.Set
	_, _ = hex.Decode(id[:], []byte(d.ID[2:]))
	if len(d.MoveSet1) == 66 {
		_, _ = hex.Decode(set1[:], []byte(d.MoveSet1[2:12]))
	}
	if len(d.MoveSet2) == 66 {
		_, _ = hex.Decode(set2[:], []byte(d.MoveSet2[2:12]))
	}
	start1, _ := strconv.ParseUint(d.StartPower1, 10, 64)
	end1, _ := strconv.ParseUint(d.EndPower1, 10, 64)
	start2, _ := strconv.ParseUint(d.StartPower2, 10, 64)
	end2, _ := strconv.ParseUint(d.EndPower2, 10, 64)
	duel := model.Duel{
		ID:        id,
		Wiz1:      wiz1,
		Wiz2:      wiz2,
		Aff1:      d.Affinity1,
		Aff2:      d.Affinity2,
		Start:     d.StartBlock,
		StartPow1: start1,
		StartPow2: start2,
		Deadline:  d.TimeoutBlock,
		EndPow1:   end1,
		EndPow2:   end2,
		End:       d.EndBlock,
		Timeout:   d.TimedOut,
		Set1:      set1,
		Set2:      set2,
	}
	return duel
}
