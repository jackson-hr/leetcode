package common

import (
	"fmt"
	"testing"
)

func TestArrayToTree(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	tree := ArrayToTree(arr)
	PrintTree(tree)
}

func TestArrayToSingleList(t *testing.T) {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7}
	list1 := ArrayToSingleList(arr1)
	fmt.Println(list1)

	var arr2 = []int{1, 2, 3, 4, 5, 6, 7, 1}
	list2 := ArrayToSingleList(arr2)
	fmt.Println(list2)

	var arr3 = []int{1, 2, 3, 4, 5, 6, 7, 4}
	list3 := ArrayToSingleList(arr3)
	fmt.Println(list3)
}
