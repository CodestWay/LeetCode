package easy

func licenseKeyFormatting(s string, k int) string {
	res := ""
	index := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			c := s[i]
			if c <= 'z' && c >= 'a' {
				c -= 32
			}
			if index == k {
				res = string(c) + "-" + res
				index = 0
			} else {
				res = string(c) + res
			}
			index++
		}
	}
	return res
}
