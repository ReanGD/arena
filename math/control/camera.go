package control

import (
	"math"

	"github.com/ReanGD/arena/math/la"
)

type Camera struct {
	matrix   la.Mat3
	center   la.Vec2
	position la.Vec2
	rotRad   float64
	zoom     float64
	zoomMin  float64
	zoomMax  float64
}

func NewCamera(viewPort la.Rect, zoomMin, zoomMax float64) *Camera {
	res := &Camera{
		matrix:   la.Mat3One,
		center:   viewPort.CenterVec2(),
		position: la.Vec2Zero,
		rotRad:   0.0,
		zoom:     1.0,
		zoomMin:  zoomMin,
		zoomMax:  zoomMax,
	}

	res.update()
	return res
}

func (c *Camera) update() {
	c.matrix = la.NewTranslateVec(c.position.Neg().SubVec(c.center)).
		ScaleAll(c.zoom).
		Rotate(c.rotRad).
		TranslateVec(c.center)
}

func (c *Camera) GetMatrix() la.Mat3 {
	return c.matrix
}

func (c *Camera) ScreenToWorld(screenPos la.Vec2) (la.Vec2, bool) {
	if c.matrix.IsInvertible() {
		return c.matrix.Inverse().MulVec(screenPos), true
	} else {
		return la.Vec2Zero, false
	}
}

func (c *Camera) Zoom(dt float64) {
	c.zoom = la.Clamp(c.zoom*math.Pow(1.01, dt), c.zoomMin, c.zoomMax)
	c.update()
}

func (c *Camera) Move(dx, dy float64) {
	c.position = c.position.Add(dx, dy)
	c.update()
}

func (c *Camera) MoveX(dx float64) {
	c.position[0] += dx
	c.update()
}

func (c *Camera) MoveY(dy float64) {
	c.position[1] += dy
	c.update()
}

func (c *Camera) RotateRad(theta float64) {
	c.rotRad += theta
	c.update()
}

func (c *Camera) RotateDeg(theta float64) {
	c.rotRad += la.DegToRad(theta)
	c.update()
}

func (c *Camera) Reset() {
	c.position = la.Vec2Zero
	c.rotRad = 0.0
	c.zoom = 0.0

	c.update()
}
