package popcount

func PopcountLoop(x uint64) uint64 {
	var n uint64 = 0
	for x != 0 {
		n += x & 1
		x >>= 1
	}
	return n
}
