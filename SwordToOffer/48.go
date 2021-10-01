package SwordToOffer

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	dict := map[byte]bool{}
	str := []byte(s)
	left := 0
	for right, val := range str {
		for dict[val] {
			_, ok := dict[str[left]]
			if ok {
				dict[str[left]] = false
			}
			left++
		}
		dict[val] = true
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}
