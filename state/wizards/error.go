// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package wizards

import "github.com/pkg/errors"

func isVMErr(err error) bool {
	if err == nil {
		return false
	}
	err = errors.Cause(err)
	return err.Error() == "VM execution error."
}
