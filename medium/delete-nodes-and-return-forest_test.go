package medium

import (
	"fmt"
	"github.com/jackson-hr/leetcode-go/common"
	"testing"
)

var delCase = []common.SliceAndUptSlice{
	{Element: []int{1, 2, 3, 4, 5, 6, 7}, Upt: []int{3, 5}},
}

func TestDelNodes(t *testing.T) {
	for _, c := range delCase {
		tree := common.ArrayToTree(c.Element)
		fmt.Println("del before:")
		common.PrintTree(tree)
		list := delNodes(tree, c.Upt)
		fmt.Println("\ndel after:")
		for _, t := range list {
			common.PrintTree(t)
		}
		fmt.Println("")
	}
}
