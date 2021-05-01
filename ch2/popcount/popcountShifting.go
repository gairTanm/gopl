package popcount

func PopcountShifting(x uint64) int {
	n := 0
	for i := uint64(0); i < 64; i++ {
		if x&(i<<1) != 0 {
			n++
		}
	}
	return n
}
