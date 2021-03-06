package main

import (
	"fmt"

	"gopl/ch6/geometry"
)

func main() {
	var p, q geometry.Point
	p.X, p.Y = 1.0, 1.0
	q.X, q.Y = 1.0, 1.0
	perim := geometry.Path{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
	}
	var scale float64 = 2.0
	fmt.Printf("%s\nLength of path: %.5f\n", perim, perim.Distance())
	fmt.Println("The distance between", p, "and", q, "is", p.Distance(&q))
	p.Scale(scale)
	fmt.Println("After scaling by", scale, ", p is:", p)
	scalePointBy := p.Scale
	scalePointBy(scale)
	fmt.Println(p)
	scalePointBy(scale)
	fmt.Println(p)
	scalePointBy(scale)
	fmt.Println(p)
}
