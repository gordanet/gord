package utils

import (
	"fmt"

	"github.com/gordanet/gord/domain/consensus/utils/constants"
)

// FormatGor takes the amount of sompis as uint64, and returns amount of GOR with 8  decimal places
func FormatGor(amount uint64) string {
	res := "                   "
	if amount > 0 {
		res = fmt.Sprintf("%19.8f", float64(amount)/constants.SompiPerGor)
	}
	return res
}
