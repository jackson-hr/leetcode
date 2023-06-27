package easy

import "github.com/jackson-hr/leetcode-go/common"

func hasCycle(head *common.ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next
		if head == fast {
			return true
		}
	}
	return false
}
