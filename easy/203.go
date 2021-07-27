package easy

import . "LeetCode/tool"

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Val: -1}
	newHead.Next = head
	temp := newHead
	for temp != nil && temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return newHead.Next
}
