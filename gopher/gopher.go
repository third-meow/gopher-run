package gopher

import (
	"github.com/veandco/go-sdl2/sdl"
	"gopher-run/loadImgTool"
)

const (
	SCAN_H = 43
	SCAN_J = 44
	SCAN_K = 45
	SCAN_L = 46
)

//addes b to a (a[n] += b[n] to all elements)
func apply(a, b *[2]int32) {
	for i := 0; i < len(*a); i += 1 {
		(*a)[i] += (*b)[i]
	}
}

type Gopher struct {
	pos, acl         [2]int32
	speed            int32
	height, width    int32
	srcRect, posRect *sdl.Rect
	tex              *sdl.Texture
}

//setup gopher
func (g *Gopher) Setup(renderer *sdl.Renderer, x, y, speed int32) {

	//set texture
	g.tex = loadImgTool.TextureFromBMP(renderer, "sprites/run.bmp")

	//set height and width
	g.height = 180
	g.width = 194

	//set source rect
	g.srcRect = &sdl.Rect{X: 0, Y: 0, W: g.width, H: g.height}

	g.pos[0] = x
	g.pos[1] = y
	g.speed = speed

	//update position rect
	g.updatePosRect()
}

//update position rectagle
func (g *Gopher) updatePosRect() {
	g.posRect = &sdl.Rect{X: g.pos[0], Y: g.pos[1], W: g.width, H: g.height}
}

//limit acceleration
func (g *Gopher) limitAcl() {
	if g.acl[0] > g.speed {
		g.acl[0] = g.speed
	} else if g.acl[0] < -g.speed {
		g.acl[0] = -g.speed
	}

	if g.acl[1] > g.speed {
		g.acl[1] = g.speed
	} else if g.acl[1] < -g.speed {
		g.acl[1] = -g.speed
	}
}

//update position, etc
func (g *Gopher) Update(keyboardState []uint8) {
	//update accel based on keyboard state
	if keyboardState[SCAN_H] == 1 {
		g.acl[0] -= 1
	}
	if keyboardState[SCAN_J] == 1 {
		g.acl[1] -= 1
	}
	if keyboardState[SCAN_K] == 1 {
		g.acl[1] += 1
	}
	if keyboardState[SCAN_L] == 1 {
		g.acl[0] += 1
	}

	//limit max acceleration
	g.limitAcl()

	//update position by acl
	apply(&g.pos, &g.acl)

	//update position rectangle
	g.updatePosRect()
}

//return data for renderer to draw
func (g *Gopher) GetDrawData() (*sdl.Texture, *sdl.Rect, *sdl.Rect) {
	return g.tex, g.srcRect, g.posRect
}