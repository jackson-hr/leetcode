package easy

import (
	"fmt"
	"testing"
)

var Cases2640 = [][]int{
	[]int{1, 2, 2, 1, 1, 0},
	[]int{0, 1},
}

func TestApplyOperations(t *testing.T) {
	for _, c := range Cases2640 {
		ca := make([]int, len(c))
		copy(ca, c)
		fmt.Println(applyOperations(c), " case:", ca)
	}
}
