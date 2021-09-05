package SwordToOffer

import . "LeetCode/tool"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return jude2(root.Left, root.Right)
}

func jude2(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	if left.Left == nil && right.Right == nil {
		if left.Right == nil && right.Left == nil {
			return true
		}
		if left.Right == nil || right.Left == nil {
			return false
		}
		return jude2(left.Right, right.Left)
	}

	if left.Left == nil || right.Right == nil {
		return false
	}

	if left.Right == nil && right.Left == nil {
		return jude2(left.Left, right.Right)
	}
	if left.Right == nil || right.Left == nil {
		return false
	}
	return jude2(left.Right, right.Left) && jude2(left.Left, right.Right)
}
