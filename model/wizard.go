// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package model

import (
	"math/rand"
	"sort"
)

type Wizard struct {
	ID     uint64
	Owner  string
	Player string
	Aff    uint8
	Pow    uint64
	Ready  bool
	Nonce  uint32
}

type WizardFilter func(Wizard) bool

type WizardSort func(Wizard, Wizard) bool

type WizardList []Wizard

func (wl WizardList) Filter(filter WizardFilter) WizardList {
	var dup WizardList
	for _, w := range wl {
		if !filter(w) {
			continue
		}
		dup = append(dup, w)
	}
	return dup
}

func (wl WizardList) Sort(less WizardSort) WizardList {
	sort.Slice(wl, func(i int, j int) bool {
		return less(wl[i], wl[j])
	})
	return wl
}

func (wl WizardList) IDs() []uint64 {
	ids := make([]uint64, 0, len(wl))
	for _, w := range wl {
		ids = append(ids, w.ID)
	}
	return ids
}

func (wl WizardList) Count() uint {
	return uint(len(wl))
}

func (wl WizardList) Sample(count uint) WizardList {
	dup := wl.Random()
	if uint(len(dup)) > count {
		dup = dup[:count]
	}
	return dup
}

func (wl WizardList) Random() WizardList {
	dup := make(WizardList, 0, len(wl))
	n := len(wl)
	for i := 0; i < n; i++ {
		j := rand.Intn(len(wl))
		dup = append(dup, wl[j])
		wl[j], wl[len(wl)-1] = wl[len(wl)-1], wl[j]
		wl = wl[0 : len(wl)-1]
	}
	return dup
}
