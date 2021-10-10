package easy

func arrangeCoins(n int) int {
	sum := 0
	curr := 0
	for {
		if sum == n {
			return curr
		}
		if sum > n {
			return curr - 1
		}
		curr += 1
		sum += curr
	}
}
