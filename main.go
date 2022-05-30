package main

import (
	"image"
	_ "image/png"
	"os"

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

	boySprite, err := getSPrite("boy.png")
	if err != nil {
		panic(err)
	}
	boyPosition := win.Bounds().Center()

	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))
		boySprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(boyPosition))
		win.Update()
	}
}

func getSPrite(path string) (*pixel.Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	pic := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(pic, pic.Bounds()), nil
}
