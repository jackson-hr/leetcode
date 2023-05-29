package leetcode

import (
	"fmt"
	"testing"
)

var caseExample = [][][]int{
	{{0, 1}, {1, 0}},
	{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
	{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
	{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}},
	{{0, 1, 1, 0, 0, 0}, {0, 1, 0, 1, 1, 0}, {0, 1, 1, 0, 1, 0}, {0, 0, 0, 1, 1, 0}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 0}},
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	for _, v := range caseExample {
		res := ShortestPathBinaryMatrix(v)
		fmt.Println(v, " res=", res)
	}
}
