package layout

type Point struct {
	X, Y float64
}

type Rectangle struct {
	Min, Max Point
}

// Rect returns a rectangle with the given coordinates.
// The coordinates do not need to be ordered.
// as same as in the standard library, image.Rect
func Rect(x1, y1, x2, y2 float64) Rectangle {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return Rectangle{
		Min: Point{X: x1, Y: y1},
		Max: Point{X: x2, Y: y2},
	}
}

func (r Rectangle) Dx() float64 {
	return r.Max.X - r.Min.X
}

func (r Rectangle) Dy() float64 {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle) Cx() float64 {
	return (r.Min.X + r.Max.X) / 2
}

func (r Rectangle) Cy() float64 {
	return (r.Min.Y + r.Max.Y) / 2
}
