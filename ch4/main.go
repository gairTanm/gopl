package main

import (
	"fmt"
	"gopl/ch4/funcs"
)

func main() {
	s := []string{"lorem", "ipsum", "ipsum", "ipsum", "takeshi"}
	s = funcs.InPlaceDup(s)
	fmt.Println(s)
}
