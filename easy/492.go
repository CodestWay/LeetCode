package easy

import "math"

func constructRectangle(area int) []int {
	for i := int(math.Sqrt(float64(area))); i > 0; i-- {
		if area/i*i == area {
			return []int{area / i, i}
		}
	}
	return []int{}
}
