package tools

import (
	"fmt"
	"math"
)

// description int -> uint8
func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	panic(fmt.Sprintf("%d is out of the uint8 range", n))
}

// @description float64 -> int
func IntFromFloat64(n float64) int {
	if math.MinInt32 <= n && n <= math.MaxInt32 {
		whole, fraction := math.Modf(n)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", n))
}
