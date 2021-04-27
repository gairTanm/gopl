package main

//usage, go run echo.go '<sep>' <args>
//eg, go run echo.go '&' tanmay hey
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Called by,", os.Args[0])
	fmt.Println(strings.Join(os.Args[2:], os.Args[1]))
}
