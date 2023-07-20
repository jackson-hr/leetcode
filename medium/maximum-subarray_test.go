package medium

import (
	"fmt"
	"testing"
)

var case53 = []struct {
	c []int
	r int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{5, 4, -1, 7, 8}, 23},
	{[]int{1}, 1},
}

func TestMaxSubArray(t *testing.T) {
	for _, s := range case53 {
		if res := maxSubArray(s.c); res != s.r {
			fmt.Println("case:", s.c, " res:", res, " should:", s.r)
			t.Fail()
		}
	}
}
