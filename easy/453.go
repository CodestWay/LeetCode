package easy

func minMoves(nums []int) int {
	min := nums[0]
	sum := 0
	n := len(nums)
	for _, num := range nums {
		if num < min {
			min = num
		}
		sum += num
	}
	return sum - n*min
}
