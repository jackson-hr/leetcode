package easy

import (
	"fmt"
	"testing"
)

var case88 = []struct {
	m  int
	n  int
	n1 []int
	n2 []int
}{
	{m: 3, n: 3, n1: []int{1, 2, 3, 0, 0, 0}, n2: []int{2, 5, 6}},
	{m: 1, n: 0, n1: []int{1}, n2: []int{}},
	{m: 0, n: 1, n1: []int{0}, n2: []int{1}},
}

func TestMerge(t *testing.T) {
	for _, s := range case88 {
		s := s
		merge(s.n1, s.m, s.n2, s.n)
		fmt.Println(s.n1)
	}
}

func TestMerge1(t *testing.T) {
	for _, s := range case88 {
		s := s
		merge1(s.n1, s.m, s.n2, s.n)
		fmt.Println(s.n1)
	}
}

func TestMerge2(t *testing.T) {
	for _, s := range case88 {
		s := s
		merge2(s.n1, s.m, s.n2, s.n)
		fmt.Println(s.n1)
	}
}
