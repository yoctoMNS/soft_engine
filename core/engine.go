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
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		sdl.Log("Failed to initialize SDL: %s", sdl.GetError())
		return false
	}

	if err := img.Init(img.INIT_JPG | img.INIT_PNG); err != nil {
		sdl.Log("Failed to initialize SDL: %s", sdl.GetError())
		return false
	}

	window, err := sdl.CreateWindow(
		"Soft Engine",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
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

	GetTextureManagerInstance().Load("tree", "assets/tree.png")

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
}

func (e *engine) Render() {
	e.Renderer.SetDrawColor(255, 0, 0, 255)
	e.Renderer.Clear()

	GetTextureManagerInstance().Draw("tree", 100, 100, 347, 465, sdl.FLIP_NONE)

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

func GetEngineInstance() *engine {
	if e == nil {
		e = &engine{}
	}

	return e
}
