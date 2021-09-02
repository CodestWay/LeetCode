package SwordToOffer

import . "LeetCode/tool"

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	flag := true
	res := make([][]int, 0)
	tmp1 := make([]*TreeNode, 0)
	tmp1 = append(tmp1, root)
	for len(tmp1) != 0 {
		tmp2 := make([]*TreeNode, 0)
		tmp3 := make([]int, 0)
		if flag {
			for i := 0; i < len(tmp1); i++ {
				node := tmp1[i]
				tmp3 = append(tmp3, node.Val)
				if node.Left != nil {
					tmp2 = append(tmp2, node.Left)
				}
				if node.Right != nil {
					tmp2 = append(tmp2, node.Right)
				}
			}
		} else {
			for i := len(tmp1) - 1; i >= 0; i-- {
				node := tmp1[i]
				tmp3 = append(tmp3, node.Val)
				if node.Right != nil {
					tmp := make([]*TreeNode, 0)
					tmp = append(tmp, node.Right)
					tmp2 = append(tmp, tmp2...)
				}
				if node.Left != nil {
					tmp := make([]*TreeNode, 0)
					tmp = append(tmp, node.Left)
					tmp2 = append(tmp, tmp2...)
				}
			}
		}

		tmp1 = tmp2
		flag = !flag
		res = append(res, tmp3)
	}
	return res
}
