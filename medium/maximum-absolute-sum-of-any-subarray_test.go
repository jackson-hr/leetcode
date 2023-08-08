package medium

import "testing"

var case1749 = []struct {
	c []int
	r int
}{
	{c: []int{1, -3, 2, 3, -4}, r: 5},
	{c: []int{2, -5, 1, -4, 3, -2}, r: 8},
}

func TestMaxAbsoluteSum(t *testing.T) {
	for _, c := range case1749 {
		if res := maxAbsoluteSum(c.c); res != c.r {
			t.Log("case:", c, " res is ", res)
			t.Fail()
		}
	}
}
