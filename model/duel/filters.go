// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package duel

import (
	"github.com/awishformore/cancoillotte/model"
)

func Not(filter model.DuelFilter) model.DuelFilter {
	return func(d model.Duel) bool {
		return !filter(d)
	}
}

func Timeout(d model.Duel) bool {
	return d.Timeout
}

func Wiz1(wizID uint64) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Wiz1 == wizID
	}
}

func Wiz2(wizID uint64) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Wiz2 == wizID
	}
}

func Owner1(owner string) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Owner1 == owner
	}
}

func Owner2(owner string) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Owner2 == owner
	}
}

func Player1(player string) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Player1 == player
	}
}

func Player2(player string) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Player2 == player
	}
}

func Aff1(aff uint8) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Aff1 == aff
	}
}

func Aff2(aff uint8) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Aff2 == aff
	}
}

func Set1Zero(d model.Duel) bool {
	return d.Set1 == model.Set{}
}

func Set2Zero(d model.Duel) bool {
	return d.Set2 == model.Set{}
}

func Within(ratio float64) model.DuelFilter {
	return func(d model.Duel) bool {
		lower := uint64(float64(d.StartPow1) / ratio)
		upper := uint64(float64(d.StartPow1) * ratio)
		return d.StartPow2 >= lower && d.StartPow2 <= upper
	}
}

func Weaker(ratio float64) model.DuelFilter {
	return func(d model.Duel) bool {
		lower := uint64(float64(d.StartPow1) / ratio)
		return d.StartPow2 < lower
	}
}

func Stronger(ratio float64) model.DuelFilter {
	return func(d model.Duel) bool {
		upper := uint64(float64(d.StartPow1) * ratio)
		return d.StartPow2 > upper
	}
}

func Expired(height uint64) model.DuelFilter {
	return func(d model.Duel) bool {
		return d.Deadline <= height
	}
}
