package easy

import "github.com/jackson-hr/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *common.TreeNode) []int {
	var ans []int
	inorder(root, &ans)
	return ans
}

func inorder(root *common.TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorder(root.Right, ans)
}
