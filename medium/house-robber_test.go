package medium

import (
	"fmt"
	"testing"
)

var robberCase = [][]int{
	{1, 2, 3, 1},
	{2, 7, 9, 3, 1},
}

func TestRob(t *testing.T) {
	for _, c := range robberCase {
		fmt.Println(rob(c), " case:", c)
	}
}

func TestRob2(t *testing.T) {
	for _, c := range robberCase {
		fmt.Println(rob2(c), " case:", c)
	}
}

func BenchmarkRob1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range robberCase {
			rob(c)
		}
	}
}

func BenchmarkRob2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range robberCase {
			rob2(c)
		}
	}
}
