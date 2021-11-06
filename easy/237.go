package easy

import . "LeetCode/tool"

func deleteNode(node *ListNode) {
	curr := node
	next := curr.Next
	nNext := next.Next
	curr.Val = next.Val
	curr.Next = nNext
}
