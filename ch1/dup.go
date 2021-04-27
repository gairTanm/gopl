package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		count[input.Text()]++
	}
	for line, cnt := range count {
		if cnt > 1 {
			fmt.Println(line, "appeared", cnt, "times")
		}
	}
}
