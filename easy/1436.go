package easy

func destCity(paths [][]string) string {
	myMap := make(map[string]struct{})
	for i := 0; i < len(paths); i++ {
		myMap[paths[i][0]] = struct{}{}
	}
	for i := 0; i < len(paths); i++ {
		_, flag := myMap[paths[i][1]]
		if !flag {
			return paths[i][1]
		}
	}
	return ""
}
