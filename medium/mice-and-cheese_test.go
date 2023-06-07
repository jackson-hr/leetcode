package medium

import (
	"fmt"
	"testing"
)

var case2611 = [][][]int{
	[][]int{
		[]int{1, 1, 3, 4},
		[]int{4, 4, 1, 1},
		[]int{2},
	},
	[][]int{
		[]int{1, 1},
		[]int{1, 1},
		[]int{2},
	},
}

func TestMiceAndCheese(t *testing.T) {
	for _, c := range case2611 {
		fmt.Println(miceAndCheese(c[0], c[1], c[2][0]), " case:", c)
	}
}
