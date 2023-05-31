package easy

import (
	"fmt"
	"testing"
)

var sqrtxCase = []int{4, 8, 9, 10, 12, 16, 1024, 625, 14884}

func TestSqrt(t *testing.T) {
	for _, c := range sqrtxCase {
		fmt.Println(mySqrt(c), " case:", c)
	}
}

func TestSqrt2(t *testing.T) {
	for _, c := range sqrtxCase {
		fmt.Println(mySqrt2(c), " case:", c)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range sqrtxCase {
			mySqrt(c)
		}
	}
}

func BenchmarkSqrt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range sqrtxCase {
			mySqrt2(c)
		}
	}
}
