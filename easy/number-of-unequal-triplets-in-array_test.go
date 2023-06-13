package easy

import (
	"fmt"
	"testing"
)

var case2475 = [][]int{
	{4, 3, 2, 2, 3},
	{4, 4, 2, 4, 3},
	{1, 1, 1, 1, 1},
}

func TestUnequalTriplets(t *testing.T) {
	for _, c := range case2475 {
		c := c
		fmt.Println(unequalTriplets(c), " case:", c)
	}
}

func TestUnequalTriplets1(t *testing.T) {
	for _, c := range case2475 {
		c := c
		fmt.Println(unequalTriplets1(c), " case:", c)
	}
}
