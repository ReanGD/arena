package sprite

import (
	"image"
	"image/color"

	"github.com/ReanGD/arena/math/la"
	"github.com/hajimehoshi/ebiten/v2"
)

type SubImage struct {
	img  *ebiten.Image
	rect la.Rect
}

func NewSubImage(img *ebiten.Image, rect la.Rect) *SubImage {
	return &SubImage{
		img:  img,
		rect: rect,
	}
}

func NewSubImageFromImage(img *ebiten.Image) *SubImage {
	rc := img.Bounds()
	return &SubImage{
		img: img,
		rect: la.NewRect(
			int32(rc.Min.X), int32(rc.Min.Y),
			int32(rc.Dx()), int32(rc.Dy()),
		),
	}
}

func (s *SubImage) ColorModel() color.Model {
	return s.img.ColorModel()
}

func (s *SubImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(s.rect.W), int(s.rect.H))
}

func (s *SubImage) At(x, y int) color.Color {
	return s.img.At(x+int(s.rect.X), y+int(s.rect.Y))
}

func (s *SubImage) Set(x, y int, c color.Color) {
	s.img.Set(x+int(s.rect.X), y+int(s.rect.Y), c)
}
