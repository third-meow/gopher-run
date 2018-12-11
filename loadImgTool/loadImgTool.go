package loadImgTool

import (
	"github.com/veandco/go-sdl2/sdl"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func TextureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	//load image
	img, err := sdl.LoadBMP(filename)
	errCheck(err)
	defer img.Free()

	//create texture
	tex, err := renderer.CreateTextureFromSurface(img)
	errCheck(err)

	return tex
}
