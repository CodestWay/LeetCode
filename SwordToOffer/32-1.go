package SwordToOffer

import . "LeetCode/tool"

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	tmp1 := make([]*TreeNode, 0)
	if root == nil {
		return []int{}
	}
	tmp1 = append(tmp1, root)
	for len(tmp1) != 0 {
		tmp2 := make([]*TreeNode, 0)
		for i := 0; i < len(tmp1); i++ {
			node := tmp1[i]
			res = append(res, node.Val)
			if node.Left != nil {
				tmp2 = append(tmp2, node.Left)
			}
			if node.Right != nil {
				tmp2 = append(tmp2, node.Right)
			}
		}
		tmp1 = tmp2
	}
	return res
}
