package gopher

import (
	"github.com/veandco/go-sdl2/sdl"
	"gopher-run/collision"
	"gopher-run/loadImgTool"
)

type Block struct {
	pos, acl            [2]float64
	imgHeight, imgWidth int
	Height, Width       int
	srcRect, posRect    *sdl.Rect
	tex                 *sdl.Texture
}

func (b *Block) Setup(renderer *sdl.Renderer, x, y int) {

	//set texture
	b.tex = loadImgTool.TextureFromBMP(renderer, "sprites/block.bmp")

	//set Height and Width
	b.imgHeight = 165
	b.Height = b.imgHeight / 2
	b.imgWidth = 165
	b.Width = b.imgWidth / 2

	//set source rect
	b.srcRect = &sdl.Rect{X: 0, Y: 0, W: int32(b.imgWidth), H: int32(b.imgHeight)}

	b.pos[0] = x
	b.pos[1] = y

	//update position rect
	b.updatePosRect()
}

//update position rectagle
func (b *Block) updatePosRect() {
	b.posRect = &sdl.Rect{X: int32(b.pos[0]), Y: int32(b.pos[1]), W: int32(b.Width), H: int32(b.Height)}
}

func (b *Block) update(col *collision.Collision) {

}

//return data for renderer to draw
func (b *Block) GetDrawData() (*sdl.Texture, *sdl.Rect, *sdl.Rect) {
	return b.tex, b.srcRect, b.posRect
}

//return x pos
func (b *Block) GetX() int {
	return int(b.pos[0])
}

//return y pos
func (b *Block) GetY() int {
	return int(b.pos[1])
}
