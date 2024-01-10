package easy

import (
	"log"
	"testing"
)

var case2660 = []struct {
	p1 []int
	p2 []int
	r  int
}{
	{p1: []int{4, 10, 7, 9}, p2: []int{6, 5, 2, 3}, r: 1},
	{p1: []int{3, 5, 7, 6}, p2: []int{8, 10, 10, 2}, r: 2},
	{p1: []int{2, 3}, p2: []int{4, 1}, r: 0},
	{p1: []int{3, 6, 10, 8}, p2: []int{9, 9, 9, 9}, r: 2},
	{p1: []int{9, 7, 10, 7}, p2: []int{10, 2, 4, 10}, r: 1},
}

func TestIsWinner(t *testing.T) {
	for _, item := range case2660 {
		res := isWinner(item.p1, item.p2)
		log.Println(item, " res:", res)
		if res != item.r {
			t.Fail()
		}
	}
}
