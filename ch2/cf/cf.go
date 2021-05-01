package cf

import (
	"flag"
	"fmt"
)

var choice = flag.String("c", "temp", "Call temp for temperature conversion and len for length conversion")

func Conversion() {
	flag.Parse()
	fmt.Println("Flag choice", *choice)
	f := *choice
	switch f {
	case "temp":
		cfTemp()
	case "len":
		cfLen()
	case "weight":
		cfWeight()
	default:
		fmt.Println("wip")
	}
}
