package easy

import (
	"log"
	"testing"

	"github.com/jackson-hr/leetcode-go/common"
)

var case141 []struct {
	Arr []int
	Res bool
} = []struct {
	Arr []int
	Res bool
}{
	{Arr: []int{3, 2, 0, -4, 2}, Res: true},
	{Arr: []int{1, 2, 1}, Res: true},
	{Arr: []int{1}, Res: false},
}

func TestHasCycle(t *testing.T) {
	for _, v := range case141 {
		v := v
		if res := hasCycle(common.ArrayToCycleList(v.Arr)); v.Res != res {
			log.Println("case:", v.Arr, v.Res, res)
			t.Fail()
		}
	}
}
