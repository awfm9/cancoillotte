// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package model

import (
	"math/rand"
	"sort"
)

type Strategy struct {
	Score float64
	Sets  []Set
}

type StrategySort func(Strategy, Strategy) bool

type StrategyList []Strategy

func (sl StrategyList) Random() StrategyList {
	dup := make(StrategyList, 0, len(sl))
	n := len(sl)
	for i := 0; i < n; i++ {
		j := rand.Intn(len(sl))
		dup = append(dup, sl[j])
		sl[j], sl[len(sl)-1] = sl[len(sl)-1], sl[j]
		sl = sl[0 : len(sl)-1]
	}
	return dup
}

func (sl StrategyList) Sort(less StrategySort) StrategyList {
	sort.Slice(sl, func(i int, j int) bool {
		return less(sl[i], sl[j])
	})
	return sl
}
