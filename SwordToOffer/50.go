package SwordToOffer

import "time"

func firstUniqChar(s string) byte {
	myMap := make(map[byte]time.Time)
	count := make(map[byte]int)
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		count[bytes[i]]++
		val := count[bytes[i]]
		if val >= 2 {
			delete(myMap, bytes[i])
		} else {
			myMap[bytes[i]] = time.Now()
		}
	}
	min := time.Now()
	res := byte(' ')
	for key, val := range myMap {
		if min.After(val) {
			min = val
			res = key
		}
	}
	return res
}
