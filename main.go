package main

import (
	"fmt"
	"log"

	"github.com/ReanGD/arena/opacha"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game interface {
	Init(screen *ebiten.Image) error
	Update(frameNum int) error
	Draw(frameNum int, screen *ebiten.Image)
}

type Arena struct {
	frameNum int
	game     Game
	initErr  error
}

func NewArena(game Game) *Arena {
	return &Arena{
		frameNum: -1,
		game:     game,
		initErr:  nil,
	}
}

func (a *Arena) Update() error {
	a.frameNum++
	if a.frameNum != 0 && a.initErr == nil {
		return a.game.Update(a.frameNum)
	}

	return a.initErr
}

func (a *Arena) Draw(screen *ebiten.Image) {
	if a.frameNum == 0 {
		a.initErr = a.game.Init(screen)
		if a.initErr == nil {
			a.initErr = a.game.Update(a.frameNum)
		}
	}

	a.game.Draw(a.frameNum, screen)

	fps := ebiten.CurrentFPS()
	tps := ebiten.CurrentTPS()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f TPS: %0.2f", fps, tps))
}

func (a *Arena) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 1024
}

func main() {
	ebiten.SetWindowSize(1280, 1024)
	ebiten.SetWindowTitle("Arena")

	game, err := opacha.NewGame()
	if err != nil {
		log.Fatal(err)
		return
	}

	arena := NewArena(game)
	if err := ebiten.RunGame(arena); err != nil {
		log.Fatal(err)
	}
}
