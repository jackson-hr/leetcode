package medium

import "github.com/jackson-hr/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *common.TreeNode) int {
	var ans int
	var dfs func(*common.TreeNode) (int, int)
	dfs = func(root *common.TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		coinsL, nodesL := dfs(root.Left)
		coinsR, nodesR := dfs(root.Right)
		coins := coinsL + coinsR + root.Val
		nodes := nodesL + nodesR + 1
		ans += abs(coins - nodes)
		return coins, nodes
	}
	dfs(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
