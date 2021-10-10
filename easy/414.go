package easy

import "sort"

func thirdMax(nums []int) int {
	myMap := make(map[int]int)
	for _, num := range nums {
		myMap[num] = 1
	}
	myList := make([]int, 0)
	for key, _ := range myMap {
		myList = append(myList, key)
	}
	sort.Ints(myList)
	if len(myList) < 3 {
		return myList[len(myList)-1]
	}
	return myList[len(myList)-3]
}
