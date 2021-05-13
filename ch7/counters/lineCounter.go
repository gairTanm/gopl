package counters

import "fmt"

type LineCounter struct {
	lines int
}

func (l *LineCounter) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			l.lines++
		}
	}
	return len(p), nil
}

func (l *LineCounter) N() int {
	return l.lines
}

func (l *LineCounter) String() string {
	return fmt.Sprintf("%d lines", l.lines)
}
