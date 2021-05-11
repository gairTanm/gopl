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
	fmt.Printf("%s\nLength of path: %.5f\n", perim, perim.Distance())
	fmt.Println("The distance between", p, "and", q, "is", p.Distance(q))
}
