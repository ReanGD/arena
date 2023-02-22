package la

import "math"

const (
	epsilon     = 1e-10
	minNormal   = float64(1.1754943508222875e-38) // 1 / 2**(127 - 1)
	piDiv180    = 0.0174532925199432957692369076848861271344287188854172545609719144
	piDiv180Inv = 57.295779513082320876798154814105170332405472466564321549160243861
)

func Equal(a, b float64) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)
	if a*b == 0 || diff < minNormal {
		return diff < epsilon*epsilon
	}

	return diff/(math.Abs(a)+math.Abs(b)) < epsilon
}

func RadToDeg(rad float64) float64 {
	return rad * piDiv180Inv
}

func DegToRad(deg float64) float64 {
	return deg * piDiv180
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	}
	return x
}
