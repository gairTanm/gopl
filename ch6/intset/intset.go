package intset

import (
	"bytes"
	"fmt"
)

//type declaration
type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, x%64
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//add an integer to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, x%64
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

func popCount(x uint64) int {
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
	word, bit := x/64, x%64
	s.words[word] &^= 1 << bit
}

//return a copy of the set
func (s *IntSet) Copy() *IntSet {
	c := &IntSet{}
	c.words = make([]uint64, len(s.words))
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
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*idx+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}
