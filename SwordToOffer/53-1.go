package SwordToOffer

func search(nums []int, target int) int {
	myMap := make(map[int]int)
	for _, val := range nums {
		myMap[val] = myMap[val] + 1
	}
	return myMap[target]
}
