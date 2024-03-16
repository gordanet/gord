// Copyright (c) 2013, 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util

import (
	"github.com/gordanet/gord/domain/consensus/utils/constants"
	"github.com/pkg/errors"
	"math"
	"strconv"
)

// AmountUnit describes a method of converting an Amount to something
// other than the base unit of a gor. The value of the AmountUnit
// is the exponent component of the decadic multiple to convert from
// an amount in gor to an amount counted in units.
type AmountUnit int

// These constants define various units used when describing a gor
// monetary amount.
const (
	AmountMegaGOR  AmountUnit = 6
	AmountKiloGOR  AmountUnit = 3
	AmountGOR      AmountUnit = 0
	AmountMilliGOR AmountUnit = -3
	AmountMicroGOR AmountUnit = -6
	AmountSompi     AmountUnit = -8
)

// String returns the unit as a string. For recognized units, the SI
// prefix is used, or "Seep" for the base unit. For all unrecognized
// units, "1eN GOR" is returned, where N is the AmountUnit.
func (u AmountUnit) String() string {
	switch u {
	case AmountMegaGOR:
		return "MGOR"
	case AmountKiloGOR:
		return "kGOR"
	case AmountGOR:
		return "GOR"
	case AmountMilliGOR:
		return "mGOR"
	case AmountMicroGOR:
		return "μGOR"
	case AmountSompi:
		return "Sompi"
	default:
		return "1e" + strconv.FormatInt(int64(u), 10) + " GOR"
	}
}

// Amount represents the base gor monetary unit (colloquially referred
// to as a `Seep'). A single Amount is equal to 1e-8 of a gor.
type Amount uint64

// round converts a floating point number, which may or may not be representable
// as an integer, to the Amount integer type by rounding to the nearest integer.
// This is performed by adding or subtracting 0.5 depending on the sign, and
// relying on integer truncation to round the value to the nearest Amount.
func round(f float64) Amount {
	if f < 0 {
		return Amount(f - 0.5)
	}
	return Amount(f + 0.5)
}

// NewAmount creates an Amount from a floating point value representing
// some value in gor. NewAmount errors if f is NaN or +-Infinity, but
// does not check that the amount is within the total amount of gor
// producible as f may not refer to an amount at a single moment in time.
//
// NewAmount is for specifically for converting GOR to Seep.
// For creating a new Amount with an int64 value which denotes a quantity of Seep,
// do a simple type conversion from type int64 to Amount.
// TODO: Refactor NewAmount. When amounts are more than 1e9 GOR, the precision
// can be higher than one seep (1e9 and 1e9+1e-8 will result as the same number)
func NewAmount(f float64) (Amount, error) {
	// The amount is only considered invalid if it cannot be represented
	// as an integer type. This may happen if f is NaN or +-Infinity.
	switch {
	case math.IsNaN(f):
		fallthrough
	case math.IsInf(f, 1):
		fallthrough
	case math.IsInf(f, -1):
		return 0, errors.New("invalid GOR amount")
	}

	return round(f * constants.SompiPerGor), nil
}

// ToUnit converts a monetary amount counted in gor base units to a
// floating point value representing an amount of gor.
func (a Amount) ToUnit(u AmountUnit) float64 {
	return float64(a) / math.Pow10(int(u+8))
}

// ToGOR is the equivalent of calling ToUnit with AmountGOR.
func (a Amount) ToGOR() float64 {
	return a.ToUnit(AmountGOR)
}

// Format formats a monetary amount counted in gor base units as a
// string for a given unit. The conversion will succeed for any unit,
// however, known units will be formated with an appended label describing
// the units with SI notation, or "Seep" for the base unit.
func (a Amount) Format(u AmountUnit) string {
	units := " " + u.String()
	return strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+8), 64) + units
}

// String is the equivalent of calling Format with AmountGOR.
func (a Amount) String() string {
	return a.Format(AmountGOR)
}

// MulF64 multiplies an Amount by a floating point value. While this is not
// an operation that must typically be done by a full node or wallet, it is
// useful for services that build on top of gor (for example, calculating
// a fee by multiplying by a percentage).
func (a Amount) MulF64(f float64) Amount {
	return round(float64(a) * f)
}
