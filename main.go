package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_STARTPOINT_X = sdl.WINDOWPOS_UNDEFINED
	SCREEN_STARTPOINT_Y = sdl.WINDOWPOS_UNDEFINED
	SCREEN_HEIGHT       = 720
	SCREEN_WIDTH        = 720
)

//renderer and window is global
var renderer *sdl.Renderer
var window *sdl.Window

//checks and handles errors
func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func updateAll() {
	return
}

//draws everything
func drawAll() {
	//draw background
	renderer.SetDrawColor(0, 5, 0, 30)
	renderer.Clear()

	//present changes
	renderer.Present()
}

//iniyializes everything
func initializeAll() {
	//initialize sdl
	err := sdl.Init(sdl.INIT_EVERYTHING)
	errCheck(err)

	//create window
	window, err = sdl.CreateWindow(
		"A Game I Guess",
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
	initializeAll()
	defer destroyAll()

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
