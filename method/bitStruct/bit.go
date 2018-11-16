package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Println(word, bit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] = 1<<bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) string() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	s := IntSet{[]uint64{1, 30, 50, 60}}
	t := IntSet{[]uint64{10, 20, 30, 60, 90, 100}}
	fmt.Println(s)
	fmt.Println(t)
	x := 30
	fmt.Println(s.Has(x))
	s.Add(90)
	fmt.Println(s)
	s.UnionWith(&t)
	fmt.Println(s)
	p := IntSet{[]uint64{1, 9, 42, 144}}
	fmt.Println(p.string() )
}
