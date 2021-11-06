package hard

func trap(height []int) int {
	leftMax := height[0]
	rightMax := height[len(height)-1]
	res := 0
	lIndex := 0
	rIndex := len(height) - 1
	for lIndex < rIndex {
		if height[lIndex] > height[rIndex] {
			rIndex--
			if height[rIndex] < min(leftMax, rightMax) {
				res += min(leftMax, rightMax) - height[rIndex]
			}
			rightMax = max(rightMax, height[rIndex])
		} else {
			lIndex++
			if height[lIndex] < max(leftMax, rightMax) {
				res += min(leftMax, rightMax) - height[lIndex]
			}
			leftMax = max(leftMax, height[lIndex])
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
