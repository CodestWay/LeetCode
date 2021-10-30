package middle

func singleNumber(nums []int) []int {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, flag := myMap[nums[i]]
		if flag {
			delete(myMap, nums[i])
		} else {
			myMap[nums[i]] = 1
		}
	}
	res := make([]int, 0)
	for val, _ := range myMap {
		res = append(res, val)
	}
	return res
}
