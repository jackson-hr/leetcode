package leetcode

import (
	"fmt"
	"testing"
)

var averageCase = [][]int{
	{1, 2, 4, 7, 10},
	{1, 3, 4, 6, 10, 12, 18},
	{1, 3, 6, 10, 12, 15},
}

func TestAverageValue(t *testing.T) {
	for _, c := range averageCase {
		fmt.Println(averageValue(c), " case:", c)
	}
}
