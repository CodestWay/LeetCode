package SwordToOffer

import . "LeetCode/tool"

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	visit2(root)
	return root
}

func visit2(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	visit2(root.Left)
	visit2(root.Right)
}
