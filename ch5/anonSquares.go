package main

import "fmt"

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	fun := square()
	fn := square()
	fmt.Println(fun())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fun())
	fmt.Println(fun())
}
