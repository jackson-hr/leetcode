package easy

import "github.com/jackson-hr/leetcode-go/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	flag := map[int]bool{head.Val: true}
	cur := head
	for cur.Next != nil {
		if _, exists := flag[cur.Next.Val]; exists {
			cur.Next = cur.Next.Next
		} else {
			flag[cur.Next.Val] = true
			cur = cur.Next
		}
	}
	return head
}

func deleteDuplicates1(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
