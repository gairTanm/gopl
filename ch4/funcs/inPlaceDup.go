package funcs

func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func InPlaceDup(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = remove(s, i)
			i--
		}
	}
	return s
}
