package coloredpoint

import "image/color"

type Point struct {
	X, Y float64
}

type coloredPoint struct {
	Point
	Color color.RGBA
}
