package core

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var t *textureManager

type textureManager struct {
	TextureMap map[string]*sdl.Texture
}

func (tm *textureManager) Load(id, fileName string) bool {
	t.TextureMap = make(map[string]*sdl.Texture)

	surface, err := img.Load(fileName)
	if err != nil {
		sdl.Log("Failed to load texture: %s, %s", fileName, sdl.GetError())
		return false
	}

	texture, err := GetEngineInstance().Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		sdl.Log("Failed to create texture from surface: %s", sdl.GetError())
		return false
	}

	t.TextureMap[id] = texture
	return true
}

func (tm *textureManager) Drop(id string) {
}

func (tm *textureManager) Clean() {
}

func (tm *textureManager) Draw(id string, x, y, width, height int32, flip sdl.RendererFlip) {
	srcRect := &sdl.Rect{
		X: 0,
		Y: 0,
		W: width,
		H: height,
	}
	dstRect := &sdl.Rect{
		X: x,
		Y: y,
		W: width,
		H: height,
	}

	GetEngineInstance().Renderer.CopyEx(
		t.TextureMap[id],
		srcRect, dstRect,
		0, nil, flip)
}

func GetTextureManagerInstance() *textureManager {
	if t == nil {
		t = &textureManager{}
	}

	return t
}
