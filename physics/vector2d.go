package physics

import "fmt"

type Vector2D struct {
	X float32
	Y float32
}

func (v1 *Vector2D) Add(v2 *Vector2D) *Vector2D {
	return &Vector2D{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 *Vector2D) Sub(v2 *Vector2D) *Vector2D {
	return &Vector2D{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 *Vector2D) Mul(scalar float32) *Vector2D {
	return &Vector2D{
		X: v1.X * scalar,
		Y: v1.Y * scalar,
	}
}

func (v1 *Vector2D) Log(msg string) {
	fmt.Println(msg, "(X Y) = (", v1.X, " ", v1.Y, ")")
}
