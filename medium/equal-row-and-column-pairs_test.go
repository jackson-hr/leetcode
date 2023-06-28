package medium

import (
	"fmt"
	"testing"
)

var case2352 = [][][]int{
	{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
	{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
}

func TestEqualPairs(t *testing.T) {
	for _, c := range case2352 {
		fmt.Println(equalPairs(c), " case:", c)
	}
}
