package helper

import "math"

// Round
// @Description: 保留小数位
// @param f
// @param n
// @return float64
func Round(f float64, n int) float64 {
	var negative bool
	if f < 0 {
		negative = true
	}
	powN := math.Pow10(n)
	f = math.Trunc(math.Abs(f)*powN+0.5) / powN
	if negative {
		return -f
	}
	return f
}
