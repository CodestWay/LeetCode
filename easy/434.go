package easy

import "strings"

func countSegments(s string) int {
	split := strings.Split(s, " ")
	res := 0
	for _, str := range split {
		if len(str) != 0 {
			res++
		}
	}
	return res
}
