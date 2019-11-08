// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package model

import (
	"math/rand"
	"sort"
)

type Outcome struct {
	Set      Set
	Average  float64
	Variance float64
	Weight   float64
}

type OutcomeFilter func(Outcome) bool

type OutcomeSort func(Outcome, Outcome) bool

type OutcomeList []Outcome

func (ol OutcomeList) Filter(filter OutcomeFilter) OutcomeList {
	var dup OutcomeList
	for _, o := range ol {
		if filter(o) {
			dup = append(dup, o)
		}
	}
	return dup
}

func (ol OutcomeList) Sort(less OutcomeSort) OutcomeList {
	sort.Slice(ol, func(i int, j int) bool {
		return less(ol[i], ol[j])
	})
	return ol
}

func (ol OutcomeList) Sample(count uint) OutcomeList {
	dup := ol.Random()
	if uint(len(dup)) > count {
		dup = dup[:count]
	}
	return dup
}

func (ol OutcomeList) Random() OutcomeList {
	dup := make(OutcomeList, 0, len(ol))
	n := len(ol)
	for i := 0; i < n; i++ {
		j := rand.Intn(len(ol))
		dup = append(dup, ol[j])
		ol[j], ol[len(ol)-1] = ol[len(ol)-1], ol[j]
		ol = ol[0 : len(ol)-1]
	}
	return dup
}
