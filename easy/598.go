package easy

func maxCount(m int, n int, ops [][]int) int {
	columns := make([]int, m)
	rows := make([]int, n)
	for _, arr := range ops {
		for i := 0; i < arr[0]; i++ {
			columns[i]++
		}
		for i := 0; i < arr[1]; i++ {
			rows[i]++
		}
	}
	coMax := 1
	roMax := 1
	tmp := columns[0]
	for coMax < m && tmp == columns[coMax] {
		coMax++
	}
	tmp = rows[0]
	for roMax < n && tmp == rows[roMax] {
		roMax++
	}
	return coMax * roMax
}
