// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package outcome

import (
	"github.com/awishformore/cancoillotte/model"
)

func Reverse(less model.OutcomeSort) model.OutcomeSort {
	return func(o1 model.Outcome, o2 model.Outcome) bool {
		return less(o2, o1)
	}
}

func AverageAsc(o1 model.Outcome, o2 model.Outcome) bool {
	return o1.Average < o2.Average
}
