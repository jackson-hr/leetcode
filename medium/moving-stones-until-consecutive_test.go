package medium

import (
	"fmt"
	"testing"
)

var case1033 = [][3]int{
	{1, 2, 5},
	{4, 3, 2},
	{3, 5, 1},
}

func TestNumMovesStones(t *testing.T) {
	for i := range case1033 {
		fmt.Println(numMovesStones(case1033[i][0], case1033[i][1], case1033[i][2]), " case:", case1033[i])
	}
}
