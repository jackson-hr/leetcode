package medium

import (
	"fmt"
	"github.com/jackson-hr/leetcode-go/common"
	"testing"
)

var case979 = []struct {
	c []int
	r int
}{
	{c: []int{3, 0, 0}, r: 2},
	{c: []int{0, 3, 0}, r: 3},
	{c: []int{1, 0, 2}, r: 2},
}

func TestDistributeCoins(t *testing.T) {
	for _, s := range case979 {
		root := common.ArrayToTree(s.c)
		if res := distributeCoins(root); res != s.r {
			fmt.Println("case:", s.c, " res:", res, " should:", s.r)
			t.Fail()
		}
	}
}
