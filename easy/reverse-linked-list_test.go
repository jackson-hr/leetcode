package easy

import (
	"fmt"
	"github.com/jackson-hr/leetcode-go/common"
	"testing"
)

var case206 = [][]int{
	{1, 2, 3, 4, 5},
	{1, 2},
	{},
}

func TestReverseList(t *testing.T) {
	for _, s := range case206 {
		list := common.ArrayToSingleList(s)
		res := reverseList(list)
		fmt.Println(res)
	}
}
