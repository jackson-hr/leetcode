package leetcode

import (
	"fmt"
	"testing"
)

var parenthesesCases = []string{
	"()",
	"()[]{}",
	"(]",
	"()[]{]",
	"{",
	"((",
	"}{",
	"(()(",
}

func TestIsValid(t *testing.T) {
	for _, p := range parenthesesCases {
		fmt.Println(isValid(p), " case:", p)
	}
}

func TestIsValid2(t *testing.T) {
	for _, p := range parenthesesCases {
		fmt.Println(isValid2(p), " case:", p)
	}
}

func BenchmarkIsValid1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, p := range parenthesesCases {
			isValid(p)
		}
	}
}

func BenchmarkIsValid2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, p := range parenthesesCases {
			isValid2(p)
		}
	}
}
