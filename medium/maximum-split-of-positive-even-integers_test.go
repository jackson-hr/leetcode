package medium

import (
	"fmt"
	"testing"
)

var case2178 = []int64{12, 7, 28}

func TestMaximumEvenSplit(t *testing.T) {
	for _, v := range case2178 {
		fmt.Println(maximumEvenSplit(v), " case", v)
	}
}

func TestMaximumEvenSplit1(t *testing.T) {
	for _, v := range case2178 {
		fmt.Println(maximumEvenSplit1(v), " case", v)
	}
}
