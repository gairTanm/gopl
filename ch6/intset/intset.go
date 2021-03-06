package intset

import (
	"bytes"
	"fmt"
)

var wordSize = 32 << (^uint(0) >> 63)

//type declaration
type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, x%wordSize
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//add an integer to the set
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, x%wordSize
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

//add a bunch of values to the set
func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}

func popCount(x uint) int {
	count := 0
	for x != 0 {
		count++
		x &= (x - 1)
	}
	return count
}

//length of the set
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount(word)
	}
	return count
}

//clear a bit set
func (s *IntSet) Clear() {
	for idx := range s.words {
		s.words[idx] = 0
	}
}

//remove a specific integer from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/wordSize, x%wordSize
	s.words[word] &^= 1 << bit
}

//return a copy of the set
func (s *IntSet) Copy() *IntSet {
	c := &IntSet{}
	c.words = make([]uint, len(s.words))
	copy(c.words, s.words)
	return c
}

//take a union of two sets
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//take an intersection of two sets
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//difference of the two sets
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//symmetric difference of the two sets
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//string method
func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for idx, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*idx+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}

//Elems() method
func (s *IntSet) Elems() []int {
	elems := make([]int, 0)
	for idx, word := range s.words {
		for j := 0; j < wordSize; j++ {
			if word&(1<<j) != 0 {
				elems = append(elems, wordSize*idx+j)
			}
		}
	}
	return elems
}
