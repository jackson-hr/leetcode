package easy

import (
	"fmt"
	"github.com/jackson-hr/leetcode-go/common"
	"testing"
)

var searchInsertCase = []common.SliceAndTarget{
	{Elements: []int{1}, Target: 1},
	{Elements: []int{1}, Target: 2},
	{Elements: []int{1}, Target: 0},
	{Elements: []int{1}, Target: -1},
	{Elements: []int{1, 3}, Target: 1},
	{Elements: []int{1, 3}, Target: 3},
	{Elements: []int{1, 3}, Target: 2},
	{Elements: []int{1, 3, 5, 6}, Target: -1},
	{Elements: []int{1, 3, 5, 6}, Target: 2},
	{Elements: []int{1, 3, 5, 6}, Target: 7},
	{Elements: []int{1, 3, 5, 6}, Target: 5},
}

func TestSearchInsertPosition(t *testing.T) {
	for _, c := range searchInsertCase {
		fmt.Println(searchInsert(c.Elements, c.Target), " slice: ", c.Elements, " val:", c.Target)
	}
}

func TestSearchInsertPosition2(t *testing.T) {
	for _, c := range searchInsertCase {
		fmt.Println(searchInsert2(c.Elements, c.Target), " slice: ", c.Elements, " val:", c.Target)
	}
}
