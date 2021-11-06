package easy

func missingNumber(nums []int) int {
	res := make([]int, len(nums)+1)
	for _, num := range nums {
		res[num] = 1
	}
	for index, val := range res {
		if val != 1 {
			return index
		}
	}
	return -1
}
