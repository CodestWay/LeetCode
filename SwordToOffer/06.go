package SwordToOffer

import . "LeetCode/tool"

var res = make([]int, 0)

func reversePrint(head *ListNode) []int {
	if head == nil {
		return res
	}
	visit(head)
	return res
}

func visit(node *ListNode) {
	if node.Next != nil {
		visit(node.Next)
	}
	res = append(res, node.Val)
}
