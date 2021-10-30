package easy

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = i
	}
	res := make([]int, 0)
out:
	for i := 0; i < len(nums1); i++ {
		val := m[nums1[i]]
		for j := val; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				continue out
			}
		}
		res = append(res, -1)
	}
	return res
}
