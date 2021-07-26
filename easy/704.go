package easy

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) < 2 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		}
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
