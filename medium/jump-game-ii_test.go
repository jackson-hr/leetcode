package medium

import (
	"fmt"
	"testing"
)

var case45 = []struct {
	c []int
	r int
}{
	{c: []int{2, 3, 1, 1, 4}, r: 2},
	{c: []int{2, 3, 0, 1, 4}, r: 2},
	{c: []int{2, 3, 1}, r: 1},
	{c: []int{1}, r: 0},
	{c: []int{0}, r: 0},
}

func TestJump(t *testing.T) {
	for _, item := range case45 {
		if res := jump(item.c); res != item.r {
			fmt.Println("case:", item, " res:", res)
			t.Fail()
		}
	}
}
