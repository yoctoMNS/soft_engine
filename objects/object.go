package objects

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/yoctoMNS/soft_engine/physics"
)

type Object interface {
	Draw()
	Update(dt float32)
	Clean()
}

type Properties struct {
	TextureID string
	W         float32
	H         float32
	Flip      sdl.RendererFlip
}

type GameObject struct {
	Properties
	Transform *physics.Transform
}
