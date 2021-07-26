package easy

func removeElement(nums []int, val int) int {
	left := 0
	right := 0
	for right < len(nums) {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
