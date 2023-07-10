package hard

import (
	"fmt"
	"testing"
)

var case42 = []struct {
	c []int
	r int
}{
	{c: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, r: 6},
	{c: []int{4, 2, 0, 3, 2, 5}, r: 9},
}

func TestTrap(t *testing.T) {
	for _, v := range case42 {
		if res := trap(v.c); res != v.r {
			fmt.Println("case:", v.c, " res:", res, " should:", v.r)
			t.Fail()
		}
	}
}
