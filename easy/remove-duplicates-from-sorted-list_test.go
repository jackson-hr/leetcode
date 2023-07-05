package easy

import (
	"fmt"
	"testing"

	"github.com/jackson-hr/leetcode-go/common"
)

var case83 = [][]int{
	{1, 1, 2},
	{1, 1, 2, 3, 3},
}

func TestDeleteDuplicates(t *testing.T) {
	for _, v := range case83 {
		list := common.ArrayToSingleList(v)
		res := deleteDuplicates(list)
		fmt.Println(res)
	}
}

func TestDeleteDuplicates1(t *testing.T) {
	for _, v := range case83 {
		list := common.ArrayToSingleList(v)
		res := deleteDuplicates1(list)
		fmt.Println(res)
	}
}
