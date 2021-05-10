package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	fmt.Printf("%s", expand(string(input), rep))
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$[^\s]+`)
	return re.ReplaceAllStringFunc(s, func(x string) string {
		return f(x[1:])
	})
}

func rep(s string) string {
	return s + "expanded"
}
