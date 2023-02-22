package la

import "math"

var Epsilon = 1e-10

func Equal(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}
