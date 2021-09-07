package SwordToOffer

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	list := make([]int, 0)
	list = append(list, 0)
	list = append(list, 1)
	for n >= 2 {
		list = append(list, (list[len(list)-1]+list[len(list)-2])%1000000007)
		n--
	}
	return list[len(list)-1]
}
