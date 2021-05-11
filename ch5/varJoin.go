package main

import "fmt"

func main() {
	s := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Println(variadicJoin(s...))
}

func variadicJoin(strs ...string) string {
	var s string
	for _, str := range strs {
		s += str
	}
	return s
}
