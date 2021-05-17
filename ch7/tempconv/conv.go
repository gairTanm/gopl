package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KtoC(k Kelvin) Celsius { return Celsius(k - 273.5) }

func KtoF(k Kelvin) Fahrenheit { return CtoF(KtoC(k)) }

func FtoK(f Fahrenheit) Kelvin { return CtoK(FtoC(f)) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FtoC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KtoC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
