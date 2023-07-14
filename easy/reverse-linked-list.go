package easy

import "github.com/jackson-hr/leetcode-go/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	var nxt *common.ListNode

	var cur = head

	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
