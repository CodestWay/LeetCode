package SwordToOffer

func findRepeatNumber(nums []int) int {
	myMap := make(map[int]int)
	for _, val := range nums {
		tmp, flag := myMap[val]
		if flag {
			return tmp
		} else {
			myMap[val] = val
		}
	}
	return -1
}
