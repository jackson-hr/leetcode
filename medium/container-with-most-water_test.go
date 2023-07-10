package medium

import (
	"fmt"
	"testing"
)

var case11 = []struct {
	c []int
	r int
}{
	{c: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, r: 49},
	{c: []int{1, 1}, r: 1},
}

func TestMaxArea(t *testing.T) {
	for _, v := range case11 {
		if res := maxArea(v.c); res != v.r {
			fmt.Println("case:", v.c, " res:", res, " should:", v.r)
			t.Fail()
		}
	}
}
