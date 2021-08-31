package SwordToOffer

func missingNumber(nums []int) int {
	if nums[0] != 0 {
		return 0
	}
	for key, val := range nums {
		if val != key {
			return val - 1
		}
	}
	return len(nums)
}
