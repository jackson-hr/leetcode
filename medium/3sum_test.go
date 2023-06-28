package medium

import (
	"log"
	"testing"
)

var case15 = [][]int{
	{-1, 0, 1, 2, -1, -4},
	{0, 1, 1},
	{0, 0, 0},
	{0, 0, 0, 0},
	{1, 2, -2, -1},
	{3, 0, -2, -1, 1, 2},
}

func TestThreeSum(t *testing.T) {
	for _, v := range case15 {
		res := threeSum(v)
		log.Println("case:", v, " res:", res)
	}
}
