package sprite

import (
	"github.com/ReanGD/arena/math/la"
	"github.com/hajimehoshi/ebiten/v2"
)

var spriteIndices = []uint16{0, 1, 2, 2, 1, 3}

type Sprite struct {
	vs [4]ebiten.Vertex
}

func NewSprite(rcTex la.Rect, width, height float32) *Sprite {
	u0 := float32(rcTex.MinX())
	v0 := float32(rcTex.MinY())
	u1 := float32(rcTex.MaxX())
	v1 := float32(rcTex.MaxY())

	return &Sprite{
		vs: [4]ebiten.Vertex{
			{
				DstX:   0.0,
				DstY:   0.0,
				SrcX:   u0,
				SrcY:   v0,
				ColorR: 1.0,
				ColorG: 1.0,
				ColorB: 1.0,
				ColorA: 1.0,
			},
			{
				DstX:   width,
				DstY:   0.0,
				SrcX:   u1,
				SrcY:   v0,
				ColorR: 1.0,
				ColorG: 1.0,
				ColorB: 1.0,
				ColorA: 1.0,
			},
			{
				DstX:   0.0,
				DstY:   height,
				SrcX:   u0,
				SrcY:   v1,
				ColorR: 1.0,
				ColorG: 1.0,
				ColorB: 1.0,
				ColorA: 1.0,
			},
			{
				DstX:   width,
				DstY:   height,
				SrcX:   u1,
				SrcY:   v1,
				ColorR: 1.0,
				ColorG: 1.0,
				ColorB: 1.0,
				ColorA: 1.0,
			},
		},
	}
}

func (s *Sprite) Draw(dst *ebiten.Image, src *ebiten.Image, projView la.Mat3, position la.Vec2) {
	op := &ebiten.DrawTrianglesOptions{
		CompositeMode: ebiten.CompositeModeSourceOver,
		Filter:        ebiten.FilterLinear,
		Address:       ebiten.AddressUnsafe,
		FillRule:      ebiten.FillAll,
	}

	vs := s.vs
	la.NewTranslateVec(position).Mul(projView).MulVertexDst(vs[:])
	dst.DrawTriangles(vs[:], spriteIndices, src, op)
}
