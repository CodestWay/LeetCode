package SwordToOffer

func reverseLeftWords(s string, n int) string {
	i := len(s)
	n %= i
	bytes := []byte(s)
	return string(bytes[n:]) + string(bytes[:n])
}
