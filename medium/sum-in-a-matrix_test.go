package medium

import (
	"fmt"
	"testing"
)

var case2679 = [][][]int{
	{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}},
	{{1}},
}

func TestMatrixSum(t *testing.T) {
	for _, v := range case2679 {
		fmt.Println(matrixSum(v))
	}
}
