package funcs

func reverse(s []string) {
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1{
		s[i], s[j] = s[j], s[i]
	}
}

func SinglePassRotate(s []string, r int) []string {
	r %= len(s)
	temp := append(s[r:], s[:r]...)
	copy(s, temp)
	return s 
}
