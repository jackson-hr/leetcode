package medium

import "github.com/jackson-hr/leetcode-go/common"

func delNodes(root *common.TreeNode, to_delete []int) []*common.TreeNode {
	var roots []*common.TreeNode
	// tag
	deleted := make(map[int]bool, len(to_delete))
	for _, v := range to_delete {
		deleted[v] = true
	}
	dfs(&roots, root, true, deleted)
	return roots
}

func dfs(roots *[]*common.TreeNode, node *common.TreeNode, isRoot bool, flags map[int]bool) *common.TreeNode {
	if node == nil {
		return nil
	}

	node.Left = dfs(roots, node.Left, flags[node.Val], flags)
	node.Right = dfs(roots, node.Right, flags[node.Val], flags)
	if flags[node.Val] {
		return nil
	}
	if isRoot {
		*roots = append(*roots, node)
	}
	return node
}
