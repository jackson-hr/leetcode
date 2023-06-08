package easy

import (
	"fmt"
	"testing"
)

var case228 = [][]int{
	[]int{0, 1, 2, 4, 5, 7},
	[]int{0, 2, 3, 4, 6, 8, 9},
}

func TestSummaryRanges(t *testing.T) {
	for _, c := range case228 {
		fmt.Println(summaryRanges(c), " case:", c)
	}
}
