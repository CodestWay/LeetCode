package middle

func majorityElement(nums []int) []int {
	count := len(nums)
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	ints := make([]int, 0)
	for val, sum := range m {
		if sum > count/3 {
			ints = append(ints, val)
		}
	}
	return ints
}
