package cf

import (
	"fmt"
	"gopl/ch2/lenconv"
	"os"
	"strconv"
)

func cfLen() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			continue
		}
		m := lenconv.Meter(l)
		f := lenconv.Feet(l)
		fmt.Println(lenconv.FtoM(f), lenconv.MtoF(m))
	}
}
