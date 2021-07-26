package mid

import (
	"math"
)

var min int = math.MaxInt32

/*
错误解答
*/
func minSubArrayLen(target int, nums []int) int {
	//visit([]int{}, nums, 0, target)
	//if min == math.MaxInt32 {
	//	return 0
	//}
	//return min
	return 1
}

//func visit(tmp []int, nums []int, index int, target int) {
//	if index >= len(nums) {
//		return
//	}
//	res := sum(tmp)
//	if res + nums[index] == target {
//		tmp = append(tmp, nums[index])
//		if len(tmp) < min {
//			min = len(tmp)
//		}
//	}
//	for i := index + 1; i < len(nums); i++ {
//		visit(tmp, nums, i, target)
//	}
//	tmp = append(tmp, nums[index])
//	for i := index + 1; i < len(nums); i++ {
//		visit(tmp, nums, i, target)
//	}
//}
//
//func sum(num []int) int {
//	res := 0
//	for _, val := range num {
//		res += val
//	}
//	return res
//}
