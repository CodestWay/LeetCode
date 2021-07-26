package easy

func sortedSquares(nums []int) []int {
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = nums[i] * nums[i]
	//}
	//sort.Ints(nums)
	//return nums
	if len(nums) == 0 {
		return nums
	}
	if len(nums) == 1 {
		nums[0] = nums[0] * nums[0]
		return nums
	}
	if nums[0] > 0 {
		for i := 0; i < len(nums); i++ {
			nums[i] = nums[i] * nums[i]
		}
		return nums
	} else if nums[len(nums)-1] < 0 {
		var res []int
		for i := len(nums) - 1; i >= 0; i-- {
			res = append(res, nums[i]*nums[i])
		}
		return res
	}
	left := 0
	right := 1
	for !(nums[left] <= 0 && nums[right] >= 0) {
		left++
		right++
	}
	var res []int
	for left >= 0 || right < len(nums) {
		if left < 0 {
			res = append(res, nums[right]*nums[right])
			right++
		} else if right >= len(nums) {
			res = append(res, nums[left]*nums[left])
			left--
		} else {
			if nums[left]*nums[left] < nums[right]*nums[right] {
				res = append(res, nums[left]*nums[left])
				left--
			} else {
				res = append(res, nums[right]*nums[right])
				right++
			}
		}
	}
	return res
}
