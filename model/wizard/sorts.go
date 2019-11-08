// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package wizard

import (
	"math/rand"

	"github.com/awishformore/cancoillotte/model"
)

func Reverse(less model.WizardSort) model.WizardSort {
	return func(w1 model.Wizard, w2 model.Wizard) bool {
		return less(w2, w1)
	}
}

func PowerAsc(w1 model.Wizard, w2 model.Wizard) bool {
	if w1.Pow == w2.Pow {
		return rand.Intn(2) == 1
	}
	return w1.Pow < w2.Pow
}
