package medium

import (
	"fmt"
	"testing"
)

var case167 = []struct {
	c []int
	t int
	r []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{2, 3, 4}, 6, []int{1, 3}},
	{[]int{-1, 0}, -1, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, v := range case167 {
		if res := twoSum(v.c, v.t); res[0] != v.r[0] || res[1] != v.r[1] {
			fmt.Println("case:", v.c, " res:", res, " should:", v.r)
			t.Fail()
		}
	}
}
