package sprite

import (
	"github.com/ReanGD/arena/math/la"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteKey uint32

type SpriteManager struct {
	sprites []*Sprite
	atlases map[uint32]*squareSpriteAtlas
	images  map[uint32]*ebiten.Image
}

func NewSpriteManager() *SpriteManager {
	return &SpriteManager{
		sprites: make([]*Sprite, 0, 128),
		atlases: make(map[uint32]*squareSpriteAtlas),
		images:  make(map[uint32]*ebiten.Image),
	}
}

// logOfSpriteSize to real sprite size:
// 3 -> 8x8
// 4 -> 16x16
// 5 -> 32x32
// 6 -> 64x64
// 7 -> 128x128
// 8 -> 256x256
func (m *SpriteManager) MakeSquareSprite(logOfSpriteSize uint32, width, height float32) (SpriteKey, *SubImage) {
	if logOfSpriteSize > 8 {
		panic("SquareSprite: logOfSpriteSize > 8")
	}
	if logOfSpriteSize < 3 {
		panic("SquareSprite: logOfSpriteSize < 3")
	}

	var ok bool
	var atlas *squareSpriteAtlas
	spriteSize := int32(1) << int32(logOfSpriteSize)
	if atlas, ok = m.atlases[logOfSpriteSize]; !ok {
		atlas = newSquareSpriteAtlas(spriteSize, 8)
		m.atlases[logOfSpriteSize] = atlas
		m.images[logOfSpriteSize] = atlas.img
	}

	img, ok := atlas.next()
	if !ok {
		panic("SquareSprite: no more space in atlas")
	}

	sprite := NewSprite(img.rect, width, height)
	m.sprites = append(m.sprites, sprite)

	spriteNum := uint32(len(m.sprites) - 1)
	imageNum := logOfSpriteSize << 24
	return SpriteKey(imageNum + spriteNum), img
}

func (m *SpriteManager) Draw(screen *ebiten.Image, projView la.Mat3, key SpriteKey, position la.Vec2) {
	spriteNum := uint32(key) & 0x00ffffff
	imageNum := uint32(key) >> 24
	if spriteNum >= uint32(len(m.sprites)) {
		panic("SpriteManager: no sprite for key")
	}
	image, ok := m.images[imageNum]
	if !ok {
		panic("SpriteManager: no image for sprite")
	}
	m.sprites[spriteNum].Draw(screen, image, projView, position)
}
