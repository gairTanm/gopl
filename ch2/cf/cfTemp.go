package cf

import (
	"fmt"
	"gopl/ch2/tempconv"
	"os"
	"strconv"
)

func cfTemp() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			continue
		}
		f := tempconv.Fahrenheit(t)
		k := tempconv.Kelvin(t)
		fmt.Println(tempconv.KtoF(k), "Fahrenheit", tempconv.FtoK(f), "Kelvin")
	}
}
