package common

import (
	"fmt"
)

// ArrayToTree array to tree based level traversal
// [1,2,3,4,5,6,7]
//
//	   1
//
//	2    3
//
// 4  5  6  7
func ArrayToTree(arr []int) *TreeNode {
	length := len(arr)
	if length < 1 {
		return nil
	}
	ind := 1
	tree := &TreeNode{Val: arr[0]}
	list := []*TreeNode{tree}
	for len(list) > 0 && ind < length {
		if list[0].Left != nil && list[0].Right != nil {
			list = list[1:]
		}
		if list[0].Left == nil {
			left := &TreeNode{Val: arr[ind]}
			list[0].Left = left
			list = append(list, left)
			ind++
		} else {
			right := &TreeNode{Val: arr[ind]}
			list[0].Right = right
			list = append(list, right)
			ind++
		}

	}
	return tree
}

// PrintTree print tree based level traversal
// [1,2,3,4,5,6,7]
//
//	   1
//
//	2    3
//
// 4  5  6  7
func PrintTree(root *TreeNode) {
	var vals []int
	nodes := []*TreeNode{nil, root}
	for len(nodes) > 0 {
		nodes = nodes[1:]
		if len(nodes) < 1 {
			break
		}
		vals = append(vals, nodes[0].Val)
		if nodes[0].Left != nil {
			nodes = append(nodes, nodes[0].Left)
		}
		if nodes[0].Right != nil {
			nodes = append(nodes, nodes[0].Right)
		}
	}
	fmt.Println(vals)
}
