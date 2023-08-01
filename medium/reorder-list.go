package medium

import (
	"github.com/jackson-hr/leetcode-go/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *common.ListNode) {
	list := make([]*common.ListNode, 0, 2)
	for node := head; node != nil; node = node.Next {
		list = append(list, node)
	}

	i, j := 0, len(list)-1
	for i < j {
		list[i].Next = list[j]
		i++
		if i == j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}
