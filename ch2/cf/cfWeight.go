package cf

import (
	"fmt"
	"gopl/ch2/cf/weightconv"
	"os"
	"strconv"
)

func cfWeight() {
	for _, arg := range os.Args[1:] {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			continue
		}
		k := weightconv.Kilogram(w)
		p := weightconv.Pound(w)
		fmt.Println(weightconv.PtoK(p), weightconv.KtoP(k))
	}
}
