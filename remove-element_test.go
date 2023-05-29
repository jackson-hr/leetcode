package leetcode

import (
	"fmt"
	"testing"
)

var removeElementCase = []SliceAndTarget{
	{Elements: []int{1}, Target: 1},
	{Elements: []int{1}, Target: 2},
	{Elements: []int{1, 1}, Target: 1},
	{Elements: []int{3, 2, 2, 3}, Target: 3},
	{Elements: []int{0, 1, 2, 2, 3, 0, 4, 2}, Target: 2},
}

func TestRemoveElement(t *testing.T) {
	for _, c := range removeElementCase {
		var s = make([]int, len(c.Elements))
		copy(s, c.Elements)
		var val = c.Target
		fmt.Println(removeElement(s, val), " slice: ", s, " val:", val)
	}
}

func TestRemoveElement2(t *testing.T) {
	for _, c := range removeElementCase {
		var s = make([]int, len(c.Elements))
		copy(s, c.Elements)
		var val = c.Target
		fmt.Println(removeElement2(s, val), " slice: ", s, " val:", val)
	}
}

func TestRemoveElement3(t *testing.T) {
	for _, c := range removeElementCase {
		var s = make([]int, len(c.Elements))
		copy(s, c.Elements)
		var val = c.Target
		fmt.Println(removeElement3(s, val), " slice: ", s, " val:", val)
	}
}
