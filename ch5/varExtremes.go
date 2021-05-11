package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var vals = []float64{1, 2, 30, -29, 201}
	min, err := variadicMin(vals...)
	must(err)
	fmt.Println("Minimum:", min)
	min, err = variadicMin()
	must(err)
	fmt.Println("Minimum:", min)
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func variadicMin(vals ...float64) (float64, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("atleast one argument is required")
	}
	var mini float64 = 1e9
	for _, val := range vals {
		mini = math.Min(mini, val)
	}
	return mini, nil
}

func variadicMax(vals ...float64) (float64, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("atleast one argument is required")
	}
	var maxi float64 = -1e9
	for _, val := range vals {
		maxi = math.Max(maxi, val)
	}
	return maxi, nil
}
