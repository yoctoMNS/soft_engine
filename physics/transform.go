package physics

import (
	"fmt"
)

type Transform struct {
	X float32
	Y float32
}

func (t *Transform) TranslateX(x float32) {
	t.X += x
}

func (t *Transform) TranslateY(y float32) {
	t.Y += y
}

func (t *Transform) Translate(v *Vector2D) {
	t.X += v.X
	t.Y += v.Y
}

func (t *Transform) Log(msg string) {
	fmt.Println(msg, "(X Y) = (", t.X, " ", t.Y, ")")
}
