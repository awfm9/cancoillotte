// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package duel

import "github.com/awishformore/cancoillotte/model"

func Reverse(sort model.DuelSort) model.DuelSort {
	return func(d1 model.Duel, d2 model.Duel) bool {
		return sort(d2, d1)
	}
}

func ByStartAsc(d1 model.Duel, d2 model.Duel) bool {
	return d1.Start > d2.Start
}
