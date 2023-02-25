package vector

import (
	"image/color"

	"github.com/eihigh/canvas"
)

// RGBA returns an alpha-premultiplied color so that c <= a.
// We silently correct the color by clipping r,g,b to a
func clipColor(col color.Color) color.RGBA {
	r, g, b, a := col.RGBA()
	if a < r {
		r = a
	}
	if a < g {
		g = a
	}
	if a < b {
		b = a
	}

	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

type Style struct {
	fillColor    color.RGBA
	strokeColor  color.RGBA
	strokeWidth  float64
	strokeCapper canvas.Capper
	strokeJoiner canvas.Joiner
	dashOffset   float64
	dashes       []float64
	colorSpace   canvas.ColorSpace
}

func NewStyle() *Style {
	return &Style{
		fillColor:    canvas.Black,
		strokeColor:  canvas.Transparent,
		strokeWidth:  1.0,
		strokeCapper: canvas.ButtCap,
		strokeJoiner: canvas.MiterJoin,
		dashOffset:   0.0,
		dashes:       []float64{},
		colorSpace:   canvas.DefaultColorSpace,
	}
}

func (s *Style) hasFill() bool {
	return s.fillColor.A != 0
}

func (s *Style) getFillColor() color.Color {
	return s.colorSpace.ToLinear(s.fillColor)
}

func (s *Style) SetFillColor(clr color.Color) {
	s.fillColor = clipColor(clr)
}

func (s *Style) hasStroke() bool {
	return s.strokeColor.A != 0 && 0.0 < s.strokeWidth
}

func (s *Style) getStrokeColor() color.Color {
	return s.colorSpace.ToLinear(s.strokeColor)
}

func (s *Style) SetStrokeColor(clr color.Color) {
	s.strokeColor = clipColor(clr)
}

func (s *Style) SetStrokeWidth(width float64) {
	s.strokeWidth = width
}

func (s *Style) SetStrokeCapper(capper canvas.Capper) {
	s.strokeCapper = capper
}

func (s *Style) SetStrokeJoiner(joiner canvas.Joiner) {
	s.strokeJoiner = joiner
}

func (s *Style) SetDashes(offset float64, dashes ...float64) {
	s.dashOffset = offset
	s.dashes = dashes
}

func (s *Style) SetColorSpace(cs canvas.ColorSpace) {
	s.colorSpace = cs
}
