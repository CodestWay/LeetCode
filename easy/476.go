package easy

import "math"

func findComplement(num int) int {
	count := 0
	tmp := num
	for num > 0 {
		num = num >> 1
		count++
	}
	pow := math.Pow(2, float64(count))
	pow -= 1
	return int(pow) - tmp
}
