package geometry

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

type ColoredPoint struct {
	Point
	Color color.RGBA
}

//euclidean distance
func (p *Point) Distance(q *Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//scale
func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

//length of the path along a given slice of points
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(&path[i])
		}
	}
	return sum
}

//string method for point
func (p Point) String() string {
	return fmt.Sprintf("%.2f, %.2f", p.X, p.Y)
}

//string method for path
func (path Path) String() string {
	var str string
	for i := range path {
		if i > 0 {
			str += path[i-1].String() + "->"
		}
	}
	str += path[len(path)-1].String()
	return str
}
