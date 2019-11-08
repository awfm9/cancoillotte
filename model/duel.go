// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package model

import (
	"math/rand"
	"sort"
)

// Duel represents a duel between two wizards.
type Duel struct {
	ID        Hash
	Wiz1      uint64
	Wiz2      uint64
	Owner1    string
	Owner2    string
	Player1   string
	Player2   string
	Aff1      uint8
	Aff2      uint8
	Start     uint64
	StartPow1 uint64
	StartPow2 uint64
	Deadline  uint64
	End       uint64
	EndPow1   uint64
	EndPow2   uint64
	Timeout   bool
	Set1      Set
	Set2      Set
}

type DuelFilter func(Duel) bool

type DuelSort func(Duel, Duel) bool

type DuelList []Duel

func (dl DuelList) Swap(swap DuelFilter) DuelList {
	for i, d := range dl {
		if swap(d) {
			dl[i].Wiz1, dl[i].Wiz2 = dl[i].Wiz2, dl[i].Wiz1
			dl[i].Owner1, dl[i].Owner2 = dl[i].Owner2, dl[i].Owner1
			dl[i].Player1, dl[i].Player2 = dl[i].Player2, dl[i].Player1
			dl[i].Aff1, dl[i].Aff2 = dl[i].Aff2, dl[i].Aff1
			dl[i].StartPow1, dl[i].StartPow2 = dl[i].StartPow2, dl[i].StartPow1
			dl[i].EndPow1, dl[i].EndPow2 = dl[i].EndPow2, dl[i].EndPow1
			dl[i].Set1, dl[i].Set2 = dl[i].Set2, dl[i].Set1
		}
	}
	return dl
}

func (dl DuelList) Filter(filter DuelFilter) DuelList {
	var dup DuelList
	for _, d := range dl {
		if filter(d) {
			dup = append(dup, d)
		}
	}
	return dup
}

func (dl DuelList) Sort(less DuelSort) DuelList {
	sort.Slice(dl, func(i int, j int) bool {
		return less(dl[i], dl[j])
	})
	return dl
}

func (dl DuelList) Count() uint {
	return uint(len(dl))
}

func (dl DuelList) Random() DuelList {
	dup := make(DuelList, 0, len(dl))
	n := len(dl)
	for i := 0; i < n; i++ {
		j := rand.Intn(len(dl))
		dup = append(dup, dl[j])
		dl[j], dl[len(dl)-1] = dl[len(dl)-1], dl[j]
		dl = dl[0 : len(dl)-1]
	}
	return dup
}

func (dl DuelList) Sample(count uint) DuelList {
	dup := dl.Random()
	if uint(len(dup)) > count {
		dup = dup[:count]
	}
	return dup
}

func (dl DuelList) Sets() []Set {
	sets := make([]Set, 0, len(dl))
	for _, d := range dl {
		sets = append(sets, d.Set2)
	}
	return sets
}
