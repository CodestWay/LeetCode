package SwordToOffer

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	mid := 0
	for left < right {
		mid = (right-left)>>2 + left
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}
