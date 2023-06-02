package medium

import (
	"fmt"
	"testing"
)

// sliceAndUptSlice double slice
type sliceAndUptSlice struct {
	Element []string
	Upt     [][]int
}

var vowelStringsCases = []sliceAndUptSlice{
	sliceAndUptSlice{Element: []string{"aba", "bcb", "ece", "aa", "e"}, Upt: [][]int{{0, 2}, {1, 4}, {1, 1}}},
	sliceAndUptSlice{Element: []string{"a", "e", "i"}, Upt: [][]int{{0, 2}, {0, 1}, {2, 2}}},
}

func TestVowelStrings(t *testing.T) {
	for _, c := range vowelStringsCases {
		fmt.Println(vowelStrings(c.Element, c.Upt))
	}
}
