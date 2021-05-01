package lenconv

func MtoF(m Meter) Feet {
	return Feet(m * 3.281)
}
func FtoM(f Feet) Meter {
	return Meter(f / 3.281)
}
