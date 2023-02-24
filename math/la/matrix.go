package la

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mat3 [9]float64

var Mat3One = Mat3{
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

func NewTranslateVec(v Vec2) Mat3 {
	return Mat3{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		v[0], v[1], 1.0,
	}
}

func NewScale(sx, sy float64) Mat3 {
	return Mat3{
		sx, 0.0, 0.0,
		0.0, sy, 0.0,
		0.0, 0.0, 1.0,
	}
}

func NewScaleAll(s float64) Mat3 {
	return Mat3{
		s, 0.0, 0.0,
		0.0, s, 0.0,
		0.0, 0.0, 1.0,
	}
}

func NewScaleVec(v Vec2) Mat3 {
	return Mat3{
		v[0], 0.0, 0.0,
		0.0, v[1], 0.0,
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

func (m Mat3) TranslateVec(v Vec2) Mat3 {
	return m.Mul(Mat3{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		v[0], v[1], 1.0,
	})
}

func (m Mat3) Scale(sx, sy float64) Mat3 {
	return m.Mul(Mat3{
		sx, 0.0, 0.0,
		0.0, sy, 0.0,
		0.0, 0.0, 1.0,
	})
}

func (m Mat3) ScaleAll(s float64) Mat3 {
	return m.Mul(Mat3{
		s, 0.0, 0.0,
		0.0, s, 0.0,
		0.0, 0.0, 1.0,
	})
}

func (m Mat3) ScaleVec(v Vec2) Mat3 {
	return m.Mul(Mat3{
		v[0], 0.0, 0.0,
		0.0, v[1], 0.0,
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

func (m Mat3) MulVec(v Vec2) Vec2 {
	return Vec2{
		m[0]*v[0] + m[3]*v[1] + m[6],
		m[1]*v[0] + m[4]*v[1] + m[7],
	}
}

func (m Mat3) MulVertexDst(v []ebiten.Vertex) {
	for i := 0; i != len(v); i++ {
		v[i].DstX, v[i].DstY = float32(m[0]*float64(v[i].DstX)+m[3]*float64(v[i].DstY)+m[6]),
			float32(m[1]*float64(v[i].DstX)+m[4]*float64(v[i].DstY)+m[7])
	}
}

func (m Mat3) IsInvertible() bool {
	return !Equal(m.Det(), 0.0)
}

func (m Mat3) Det() float64 {
	return m[0]*m[4]*m[8] + m[1]*m[5]*m[6] + m[2]*m[3]*m[7] -
		m[2]*m[4]*m[6] - m[1]*m[3]*m[8] - m[0]*m[5]*m[7]
}

func (m Mat3) Inverse() Mat3 {
	det := m.Det()
	if Equal(m.Det(), 0.0) {
		return Mat3{}
	}

	invDet := 1.0 / det
	return Mat3{
		(m[4]*m[8] - m[5]*m[7]) * invDet,
		(m[2]*m[7] - m[1]*m[8]) * invDet,
		(m[1]*m[5] - m[2]*m[4]) * invDet,
		(m[5]*m[6] - m[3]*m[8]) * invDet,
		(m[0]*m[8] - m[2]*m[6]) * invDet,
		(m[2]*m[3] - m[0]*m[5]) * invDet,
		(m[3]*m[7] - m[4]*m[6]) * invDet,
		(m[1]*m[6] - m[0]*m[7]) * invDet,
		(m[0]*m[4] - m[1]*m[3]) * invDet,
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
