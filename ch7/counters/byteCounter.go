package counters

import "io"

type ByteCounter struct {
	w       io.Writer
	written int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &ByteCounter{w, 0}
	return c, &c.written
}
