package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	errCheck(err)

	window, err := sdl.CreateWindow(
		"A Game I Guess",
		400,
		400,
		500,
		500,
		sdl.WINDOW_OPENGL)
	errCheck(err)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	errCheck(err)
	defer renderer.Destroy()

	for {
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		renderer.Present()
	}

}
