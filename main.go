package main

import (
	"os"

	"github.com/amitbikram/mmo/engine/asset"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
	cfg := pixelgl.WindowConfig{
		Title:     "MMO",
		Bounds:    pixel.R(0, 0, 1024, 768),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)
	load := asset.NewLoad(os.DirFS("./"))

	boySprite, err := load.Sprite(".boy.png")
	if err != nil {
		panic(err)
	}
	boyPosition := win.Bounds().Center()

	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))
		if win.Pressed(pixelgl.KeyLeft) {
			boyPosition.X -= 2.0
		}
		if win.Pressed(pixelgl.KeyRight) {
			boyPosition.X += 2.0
		}
		if win.Pressed(pixelgl.KeyUp) {
			boyPosition.Y += 2.0
		}
		if win.Pressed(pixelgl.KeyDown) {
			boyPosition.Y -= 2.0
		}
		boySprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 0.2).Moved(boyPosition))
		win.Update()
	}
}
