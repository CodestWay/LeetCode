package SwordToOffer

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	list := make([]int, 0)
	list = append(list, 0)
	list = append(list, 1)
	list = append(list, 2)
	for n >= 3 {
		list = append(list, (list[len(list)-1]+list[len(list)-2])%1000000007)
		n--
	}
	return list[len(list)-1]
}
