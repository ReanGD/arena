package la

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mat3 [9]float64

var One = Mat3{
	1.0, 0.0, 0.0,
	0.0, 1.0, 0.0,
	0.0, 0.0, 1.0,
}

func NewTranslate(x, y float64) Mat3 {
	return Mat3{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		x, y, 1.0,
	}
}

func NewScale(sx, sy float64) Mat3 {
	return Mat3{
		sx, 0.0, 0.0,
		0.0, sy, 0.0,
		0.0, 0.0, 1.0,
	}
}

// Theta given angle in radians
func NewRotate(theta float64) Mat3 {
	sintheta, costheta := math.Sincos(theta)
	return Mat3{
		costheta, sintheta, 0.0,
		-sintheta, costheta, 0.0,
		0.0, 0.0, 1.0,
	}
}

func (m Mat3) Translate(x, y float64) Mat3 {
	return m.Mul(Mat3{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		x, y, 1.0,
	})
}

func (m Mat3) Scale(sx, sy float64) Mat3 {
	return m.Mul(Mat3{
		sx, 0.0, 0.0,
		0.0, sy, 0.0,
		0.0, 0.0, 1.0,
	})
}

// Theta given angle in radians
func (m Mat3) Rotate(theta float64) Mat3 {
	sintheta, costheta := math.Sincos(theta)
	return m.Mul(Mat3{
		costheta, sintheta, 0.0,
		-sintheta, costheta, 0.0,
		0.0, 0.0, 1.0,
	})
}

func (m Mat3) Mul(o Mat3) Mat3 {
	return Mat3{
		m[0]*o[0] + m[1]*o[3] + m[2]*o[6],
		m[0]*o[1] + m[1]*o[4] + m[2]*o[7],
		m[0]*o[2] + m[1]*o[5] + m[2]*o[8],
		m[3]*o[0] + m[4]*o[3] + m[5]*o[6],
		m[3]*o[1] + m[4]*o[4] + m[5]*o[7],
		m[3]*o[2] + m[4]*o[5] + m[5]*o[8],
		m[6]*o[0] + m[7]*o[3] + m[8]*o[6],
		m[6]*o[1] + m[7]*o[4] + m[8]*o[7],
		m[6]*o[2] + m[7]*o[5] + m[8]*o[8],
	}
}

func (m Mat3) Apply(g *ebiten.GeoM) {
	g.SetElement(0, 0, m[0])
	g.SetElement(0, 1, m[3])
	g.SetElement(1, 0, m[1])
	g.SetElement(1, 1, m[4])
	g.SetElement(0, 2, m[6])
	g.SetElement(1, 2, m[7])
}
