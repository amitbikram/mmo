package asset

import (
	"image"
	_ "image/png"
	"io/fs"

	"github.com/faiface/pixel"
)

type Load struct {
	fileSystem fs.FS
}

func NewLoad(fileSystem fs.FS) *Load {
	return &Load{fileSystem}
}

func (load *Load) Open(path string) (fs.File, error) {
	return load.fileSystem.Open(path)
}

func (load *Load) SPrite(path string) (*pixel.Sprite, error) {
	file, err := load.Open(path)
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
