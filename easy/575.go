package easy

func distributeCandies(candyType []int) int {
	eat := len(candyType) / 2
	myMap := make(map[int]int)
	for _, val := range candyType {
		myMap[val]++
	}
	if len(myMap) > eat {
		return eat
	}
	return len(myMap)
}
