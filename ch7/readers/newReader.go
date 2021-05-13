package readers

import "io"

type newReader struct {
	s string
}

func (r *newReader) Read(p []byte) (int, error) {
	n := copy(p, r.s)
	r.s = r.s[n:]
	var err error
	if len(r.s) == 0 {
		err = io.EOF
	}
	return n, err
}

func NewReader(s string) io.Reader {
	return &newReader{s}
}
