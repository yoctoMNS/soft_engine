package core

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var e *engine

const (
	SCREEN_WIDTH  int32 = 960
	SCREEN_HEIGHT int32 = 640
)

type engine struct {
	isRunning bool
	window    *sdl.Window
	Renderer  *sdl.Renderer
}

func (e *engine) Init() bool {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		sdl.Log("Failed to initialize SDL: %s", sdl.GetError())
		return false
	}

	if err := img.Init(img.INIT_JPG | img.INIT_PNG); err != nil {
		sdl.Log("Failed to initialize SDL: %s", sdl.GetError())
		return false
	}

	window, err := sdl.CreateWindow(
		"Soft Engine",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.Log("Failed to create Window: %s", sdl.GetError())
		return false
	}
	e.window = window

	renderer, err := sdl.CreateRenderer(
		e.window,
		-1,
		sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		sdl.Log("Failed to create Renderer: %s", sdl.GetError())
		return false
	}
	e.Renderer = renderer

	e.isRunning = true
	return e.isRunning
}

func (e *engine) Clean() bool {
	return false
}

func (e *engine) Quit() {
	img.Quit()
	sdl.Quit()
	e.isRunning = false
}

func (e *engine) Update() {
	sdl.Log("test")
}

func (e *engine) Render() {
	e.Renderer.SetDrawColor(124, 218, 254, 255)
	e.Renderer.Present()
}

func (e *engine) Events() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			e.Quit()
		}
	}
}

func (e *engine) IsRunning() bool {
	return e.isRunning
}

func GetInstance() *engine {
	if e == nil {
		e = &engine{}
	}

	return e
}
