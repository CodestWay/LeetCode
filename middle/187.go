package middle

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 11 {
		return []string{}
	}
	res := make(map[string]int)
	for i := 0; i+10 <= len(s); i++ {
		res[s[i:i+10]]++
	}
	tmp := make([]string, 0)
	for str, count := range res {
		if count > 1 {
			tmp = append(tmp, str)
		}
	}
	return tmp
}
