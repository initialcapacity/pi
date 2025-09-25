package precise

import "math/big"

func DivideUint64(numerator, denominator uint64) float64 {
	bigNumerator := new(big.Int).SetUint64(numerator)
	bigDenominator := new(big.Int).SetUint64(denominator)
	ratio := new(big.Rat).SetFrac(bigNumerator, bigDenominator)

	approx, _ := ratio.Float64()
	return approx
}
