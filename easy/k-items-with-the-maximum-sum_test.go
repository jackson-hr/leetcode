package easy

import (
	"fmt"
	"testing"
)

var case2600 = [][4]int{
	{3, 2, 0, 2},
	{3, 2, 0, 4},
	{3, 1, 1, 5},
}

func TestKItemsWithMaximumSum(t *testing.T) {
	for _, v := range case2600 {
		fmt.Println(kItemsWithMaximumSum(v[0], v[1], v[2], v[3]), " case:", v)
	}
}
