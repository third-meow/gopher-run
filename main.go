package main

import (
	"github.com/veandco/go-sdl2/sdl"
	//"gopher-run/block"
	"gopher-run/gopher"
	//"gopher-run/possessed"
)

const (
	SCREEN_STARTPOINT_X = sdl.WINDOWPOS_UNDEFINED
	SCREEN_STARTPOINT_Y = sdl.WINDOWPOS_UNDEFINED
	SCREEN_HEIGHT       = 720
	SCREEN_WIDTH        = 720

	ENEMY_N = 5
	BLOCK_N = 9
)

//sdl renderer and window
var renderer *sdl.Renderer
var window *sdl.Window

//in-game entitys
var player *gopher.Gopher

//var enemies []*possessed.PossessedGopher
//var blocks []*block.Block

//checks and handles errors
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func setupEntitys() {

	player = &gopher.Gopher{}
	player.Setup(renderer, (SCREEN_WIDTH / 2), (SCREEN_HEIGHT / 2), 0.3)
	//enemies = make([]possessed.PossessedGopher, ENEMY_N)
	//blocks= make([]block.Block, BLOCK_N)

}

func updateAll() {
	player.Update(sdl.GetKeyboardState())
}

//draws everything
func drawAll() {
	//draw background
	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	//draw player
	renderer.Copy(player.GetDrawData())

	//present changes
	renderer.Present()
}

//iniyializes everything
func initializeApplication() {
	//initialize sdl
	err := sdl.Init(sdl.INIT_EVERYTHING)
	errCheck(err)

	//create window
	window, err = sdl.CreateWindow(
		"Gopher-Run",
		SCREEN_STARTPOINT_X,
		SCREEN_STARTPOINT_Y,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
		sdl.WINDOW_OPENGL)
	errCheck(err)

	//create renderer
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	errCheck(err)
}

func destroyAll() {
	window.Destroy()
	renderer.Destroy()
}

func main() {
	initializeApplication()
	defer destroyAll()

	//setup in-game entitys
	setupEntitys()

	//game loop
	for {
		updateAll()
		drawAll()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				//exit program
				return
			}
		}
	}

}
