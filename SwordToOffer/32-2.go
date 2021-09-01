package SwordToOffer

import . "LeetCode/tool"

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	tmp := make([]*TreeNode, 0)
	if root == nil {
		return [][]int{}
	}
	tmp = append(tmp, root)
	for len(tmp) != 0 {
		tmp1 := make([]*TreeNode, 0)
		tmp2 := make([]int, 0)
		for i := 0; i < len(tmp); i++ {
			node := tmp[i]
			tmp2 = append(tmp2, node.Val)
			if node.Left != nil {
				tmp1 = append(tmp1, node.Left)
			}
			if node.Right != nil {
				tmp1 = append(tmp1, node.Right)
			}
		}
		tmp = tmp1
		res = append(res, tmp2)
	}
	return res
}
