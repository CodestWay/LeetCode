package SwordToOffer

import (
	. "LeetCode/tool"
)

func copyRandomList(head *Node) *Node {
	nodeToNode := make(map[*Node]*Node)
	tmp := head
	for tmp != nil {
		nodeToNode[tmp] = &Node{
			Val: tmp.Val,
		}
		tmp = tmp.Next
	}
	tmp = head
	for tmp != nil {
		nodeToNode[tmp].Next = nodeToNode[tmp.Next]
		nodeToNode[tmp].Random = nodeToNode[tmp.Random]
		tmp = tmp.Next
	}
	return nodeToNode[head]
}
