package main

// Vec2D pos in a 2D plane
type Vec2D struct {
	X, Y int
}

func (v *Vec2D) Add(v2 Vec2D) Vec2D {
	return Vec2D{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v *Vec2D) Subst(v2 Vec2D) Vec2D {
	return Vec2D{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v *Vec2D) Mult(n int) Vec2D {
	return Vec2D{
		X: v.X * n,
		Y: v.Y * n,
	}
}
