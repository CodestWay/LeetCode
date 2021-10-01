package SwordToOffer

import . "LeetCode/tool"

func deleteNode(head *ListNode, val int) *ListNode {
	newHead := ListNode{
		Val:  0,
		Next: nil,
	}
	newHead.Next = head
	tmp := &newHead
	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		}
		tmp = tmp.Next
	}
	return newHead.Next
}
