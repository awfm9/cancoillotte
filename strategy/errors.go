// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package strategy

import "github.com/pkg/errors"

func isInDuelErr(err error) bool {
	if err == nil {
		return false
	}
	if errors.Cause(err).Error() == "Wizard currently in duel" {
		return true
	}
	return false
}

func isTooManyErr(err error) bool {
	if err == nil {
		return false
	}
	if errors.Cause(err).Error() == "Too many challenges" {
		return true
	}
	return false
}
