package middle

func longestSubsequence(arr []int, difference int) int {
	ans := 0
	f := map[int]int{}
	for _, item := range arr {
		f[item] = f[item-difference] + 1
		if f[item] > ans {
			ans = f[item]
		}
	}
	return ans
}
