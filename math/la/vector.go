package la

import (
	"fmt"
	"math"
)

type Vec2 [2]float64

var Vec2Zero = Vec2{0.0, 0.0}

func NewVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) Neg() Vec2 {
	return Vec2{-v[0], -v[1]}
}

func (v Vec2) Add(x, y float64) Vec2 {
	return Vec2{v[0] + x, v[1] + y}
}

func (v Vec2) AddVec(other Vec2) Vec2 {
	return Vec2{v[0] + other[0], v[1] + other[1]}
}

func (v Vec2) Sub(x, y float64) Vec2 {
	return Vec2{v[0] - x, v[1] - y}
}

func (v Vec2) SubVec(other Vec2) Vec2 {
	return Vec2{v[0] - other[0], v[1] - other[1]}
}

func (v Vec2) Mul(f float64) Vec2 {
	return Vec2{f * v[0], f * v[1]}
}

func (v Vec2) Div(f float64) Vec2 {
	return Vec2{v[0] / f, v[1] / f}
}

func (v Vec2) Length() float64 {
	return math.Hypot(v[0], v[1])
}

func (v Vec2) DistanceTo(other Vec2) float64 {
	return math.Hypot(v[0]-other[0], v[1]-other[1])
}

func (v Vec2) Norm(length float64) Vec2 {
	d := v.Length()
	if d == 0 {
		return Vec2{}
	}
	return Vec2{v[0] / d * length, v[1] / d * length}
}

func (v Vec2) MoveTo(other Vec2, length float64) Vec2 {
	return other.SubVec(v).Norm(length).AddVec(v)
}

func (v Vec2) Interpolate(other Vec2, t float64) Vec2 {
	return Vec2{(1.0-t)*v[0] + t*other[0], (1.0-t)*v[1] + t*other[1]}
}

func (v Vec2) Equals(other Vec2) bool {
	return Equal(v[0], other[0]) && Equal(v[1], other[1])
}

func (p Vec2) String() string {
	return fmt.Sprintf("(%g, %g)", p[0], p[1])
}
