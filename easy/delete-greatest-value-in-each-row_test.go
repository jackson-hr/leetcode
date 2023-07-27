package easy

import (
	"fmt"
	"testing"
)

var case2500 = []struct {
	c [][]int
	r int
}{
	{c: [][]int{{1, 2, 4}, {3, 3, 1}}, r: 8},
	{c: [][]int{{10}}, r: 10},
}

func TestDeleteGreatestValue(t *testing.T) {
	for _, item := range case2500 {
		if res := deleteGreatestValue(item.c); res != item.r {
			fmt.Println("case:", item, " res:", res)
			t.Fail()
		}
	}
}
