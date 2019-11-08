// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package wizard

import (
	"github.com/awishformore/cancoillotte/model"
)

func Not(filter model.WizardFilter) model.WizardFilter {
	return func(w model.Wizard) bool {
		return !filter(w)
	}
}

func Available(w model.Wizard) bool {
	return w.Ready && w.Pow > 0 && w.Player != ""
}

func ID(ids ...uint64) model.WizardFilter {
	set := make(map[uint64]struct{})
	for _, id := range ids {
		set[id] = struct{}{}
	}
	return func(w model.Wizard) bool {
		_, ok := set[w.ID]
		return ok
	}
}

func Owner(owners ...string) model.WizardFilter {
	set := make(map[string]struct{})
	for _, owner := range owners {
		set[owner] = struct{}{}
	}
	return func(w model.Wizard) bool {
		_, ok := set[w.Owner]
		return ok
	}
}

func Affinity(aff uint8) model.WizardFilter {
	return func(w model.Wizard) bool {
		return w.Aff == aff
	}
}

func Below(pow uint64) model.WizardFilter {
	return func(w model.Wizard) bool {
		return w.Pow < pow
	}
}

func Above(pow uint64) model.WizardFilter {
	return func(w model.Wizard) bool {
		return w.Pow > pow
	}
}

func Active(active func(player string, height uint64) bool, height uint64) model.WizardFilter {
	return func(w model.Wizard) bool {
		return active(w.Player, height)
	}
}
