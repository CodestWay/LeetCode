package easy

func plusOne(digits []int) []int {
	up := 0
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += up
		up = 0
		if digits[i] >= 10 {
			digits[i] %= 10
			up = 1
		}
	}
	if up > 0 {
		tmp := append(make([]int, 0), 1)
		tmp = append(tmp, digits...)
		return tmp
	}
	return digits
}
