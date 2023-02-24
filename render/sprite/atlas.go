package sprite

import (
	"github.com/ReanGD/arena/math/la"
	"github.com/hajimehoshi/ebiten/v2"
)

type squareSpriteAtlas struct {
	img           *ebiten.Image
	spritesPerRow int32
	spriteSize    int32
	nextSprite    int32
}

func newSquareSpriteAtlas(spriteSize int32, spritesPerRow int32) *squareSpriteAtlas {
	size := spriteSize * spritesPerRow
	return &squareSpriteAtlas{
		img:           ebiten.NewImage(int(size), int(size)),
		spritesPerRow: spritesPerRow,
		spriteSize:    spriteSize,
		nextSprite:    0,
	}
}

func (a *squareSpriteAtlas) next() (la.Rect, bool) {
	if a.nextSprite >= a.spritesPerRow*a.spritesPerRow {
		return la.Rect{}, false
	}

	spriteX := (a.nextSprite % a.spritesPerRow) * a.spriteSize
	spriteY := (a.nextSprite / a.spritesPerRow) * a.spriteSize
	a.nextSprite++

	return la.NewRect(
		spriteX,
		spriteY,
		a.spriteSize,
		a.spriteSize,
	), true
}
