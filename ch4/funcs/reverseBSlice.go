package funcs

func ReverseBSlice(b []byte){
	for i, j := 0, len(b); i<j; i, j = i+1, j-1{
		b[i], b[j] = b[j], b[i]
	}
}
