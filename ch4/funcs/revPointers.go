package funcs

func RevPointers(p *[4]string) {
	for i, j := 0, 3; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}
