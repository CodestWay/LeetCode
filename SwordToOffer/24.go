package SwordToOffer

import . "LeetCode/tool"

func reverseList(head *ListNode) *ListNode {
	newH := &ListNode{Val: -1}
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = nil

		next := newH.Next
		newH.Next = tmp
		tmp.Next = next
	}
	return newH.Next
}
