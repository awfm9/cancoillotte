// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package strategy

import "github.com/awishformore/cancoillotte/model"

func Reverse(less model.StrategySort) model.StrategySort {
	return func(s1 model.Strategy, s2 model.Strategy) bool {
		return less(s2, s1)
	}
}

func ScoreAsc(s1 model.Strategy, s2 model.Strategy) bool {
	return s1.Score < s2.Score
}
