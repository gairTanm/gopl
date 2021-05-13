package counters

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type WordCounter struct {
	words  int
	inWord bool
}

func leadingSpaces(p []byte) int {
	count := 0
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		if !unicode.IsSpace(r) {
			return count
		}
		cur += size
		count++
	}
	return count
}

func leadingNonSpaces(p []byte) int {
	count := 0
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		if unicode.IsSpace(r) {
			return count
		}
		cur += size
		count++
	}
	return count
}

func (w *WordCounter) Write(p []byte) (int, error) {
	cur := 0
	for {
		spaces := leadingSpaces(p[cur:])
		cur += spaces
		if spaces > 0 {
			w.inWord = false
		}
		if cur == len(p) {
			return len(p), nil
		}
		if !w.inWord {
			w.words++
		}
		w.inWord = true
		cur += leadingNonSpaces(p[cur:])
		if cur == len(p) {
			return len(p), nil
		}
	}
}

func (w *WordCounter) N() int {
	return w.words
}

func (w *WordCounter) String() string {
	return fmt.Sprintf("%d words", w.words)
}
