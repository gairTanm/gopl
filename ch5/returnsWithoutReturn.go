package main

import "fmt"

func main() {
	fmt.Println(weird())
}

func weird() (str string) {
	defer func() {
		if r := recover(); r != nil {
			str = "did it happen?"
		}
	}()
	panic("**visible panic**")
}
