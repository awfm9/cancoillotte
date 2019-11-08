// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package challenge

import (
	"github.com/awishformore/cancoillotte/model"
)

func All(model.Challenge) bool {
	return true
}

func Not(filter model.ChallengeFilter) model.ChallengeFilter {
	return func(c model.Challenge) bool {
		return !filter(c)
	}
}

func Wiz1(wiz uint64) model.ChallengeFilter {
	return func(c model.Challenge) bool {
		return c.Wiz1 == wiz
	}
}

func Wiz2(wiz uint64) model.ChallengeFilter {
	return func(c model.Challenge) bool {
		return c.Wiz2 == wiz
	}
}

func Owner1(owner string) model.ChallengeFilter {
	return func(c model.Challenge) bool {
		return c.Owner1 == owner
	}
}

func Owner2(owner string) model.ChallengeFilter {
	return func(c model.Challenge) bool {
		return c.Owner2 == owner
	}
}

func Status(status string) model.ChallengeFilter {
	return func(c model.Challenge) bool {
		return c.Status == status
	}
}
