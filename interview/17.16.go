package interview

import "math"

func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, nums[0]
	for i := 1; i < n; i++ {
		tdp0 := int(math.Max(float64(dp0), float64(dp1)))
		tdp1 := dp0 + nums[i]
		dp0 = tdp0
		dp1 = tdp1
	}
	return int(math.Max(float64(dp0), float64(dp1)))
}
