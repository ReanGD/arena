package vector

import (
	"image"
	"image/draw"
	"math"

	"github.com/ReanGD/arena/math/la"
	"github.com/eihigh/canvas"
	vimg "golang.org/x/image/vector"
)

const piDiv180 = 0.0174532925199432957692369076848861271344287188854172545609719144

func toArcCoord(center la.Vec2, r, startAngle, lenAngle float64) (la.Vec2, la.Vec2) {
	sin1, cos1 := math.Sincos(startAngle)
	sin2, cos2 := math.Sincos(startAngle + lenAngle)

	// sin(x-90) = -cos(x)
	// cos(x-90) = sin(x)
	start := center.Add(r*sin2, -r*cos2)
	end := center.Add(r*sin1, -r*cos1)
	return start, end
}

type Path struct {
	impl *canvas.Path
}

func NewPath() *Path {
	return &Path{
		impl: &canvas.Path{},
	}
}

func (p *Path) Circle(center la.Vec2, radius float64) {
	x0 := center[0] + radius
	x1 := center[0] - radius
	y := center[1]
	p.impl.MoveTo(x0, y)
	p.impl.ArcTo(radius, radius, 0.0, false, true, x1, y)
	p.impl.ArcTo(radius, radius, 0.0, false, true, x0, y)
}

// angle in radians
func (p *Path) Sector(center la.Vec2, radius, startAngle, lenAngle float64) {
	start, end := toArcCoord(center, radius, startAngle, lenAngle)
	// >= 180 degrees
	large := lenAngle >= la.Pi

	p.MoveTo(start)
	p.impl.ArcTo(radius, radius, 0, large, false, end[0], end[1])
}

// angle in radians
func (p *Path) DashSector(
	center la.Vec2, cntDash, totalDash uint, radius, startAngle, ratio float64,
) {
	lenSector := la.Pi2 / float64(totalDash)
	lenFilled := lenSector * ratio
	for i := uint(0); i != cntDash; i++ {
		p.Sector(center, radius, startAngle+lenSector*float64(i), lenFilled)
	}
}

func (p *Path) MoveTo(point la.Vec2) {
	p.impl.MoveTo(point[0], point[1])
}

func (p *Path) LineTo(point la.Vec2) {
	p.impl.LineTo(point[0], point[1])
}

func (p *Path) Line(from la.Vec2, to la.Vec2) {
	p.impl.MoveTo(from[0], from[1])
	p.impl.LineTo(to[0], to[1])
}

func (p *Path) Draw(dst draw.Image, style *Style) {
	p.DrawEx(dst, canvas.Identity, style)
}

func (p *Path) DrawEx(dst draw.Image, mat canvas.Matrix, style *Style) {
	if !style.hasFill() && !style.hasStroke() {
		return
	}

	imgSize := dst.Bounds().Size()
	pathBounds := canvas.Rect{}

	m := canvas.Identity.ReflectYAbout(float64(imgSize.Y) / 2.0).Mul(mat)

	fill := p.impl
	if style.hasFill() {
		fill = fill.Transform(m)
		if !style.hasStroke() {
			pathBounds = fill.Bounds()
		}
	}

	stroke := p.impl
	if style.hasStroke() {
		if 0 < len(style.dashes) {
			stroke = stroke.Dash(style.dashOffset, style.dashes...)
		}
		stroke = stroke.Stroke(style.strokeWidth, style.strokeCapper, style.strokeJoiner)
		stroke = stroke.Transform(m)
		pathBounds = stroke.Bounds()
	}

	dx, dy := 0, 0
	x := int(pathBounds.X)
	y := int(pathBounds.Y)
	w := int(pathBounds.W) + 1
	h := int(pathBounds.H) + 1
	if (x+w <= 0 || imgSize.X <= x) && (y+h <= 0 || imgSize.Y <= y) {
		return // outside canvas
	}

	if x < 0 {
		dx = -x
		x = 0
	}
	if y < 0 {
		dy = -y
		y = 0
	}
	if imgSize.X <= x+w {
		w = imgSize.X - x
	}
	if imgSize.Y <= y+h {
		h = imgSize.Y - y
	}
	if w <= 0 || h <= 0 {
		return // has no size
	}

	sp := image.Point{dx, dy}
	outRc := image.Rect(x, imgSize.Y-y, x+w, imgSize.Y-y-h)
	if style.hasFill() {
		ras := vimg.NewRasterizer(w, h)
		fill = fill.Translate(-float64(x), -float64(y))
		fill.ToRasterizer(ras, canvas.DPMM(1))
		clr := image.NewUniform(style.getFillColor())
		ras.Draw(dst, outRc, clr, sp)
	}
	if style.hasStroke() {
		ras := vimg.NewRasterizer(w, h)
		stroke = stroke.Translate(-float64(x), -float64(y))
		stroke.ToRasterizer(ras, canvas.DPMM(1))
		clr := image.NewUniform(style.getStrokeColor())
		ras.Draw(dst, outRc, clr, sp)
	}
}
