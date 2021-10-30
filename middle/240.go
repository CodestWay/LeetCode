package middle

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	column := len(matrix[row]) - 1
	for row <= len(matrix)-1 && column >= 0 {
		val := matrix[row][column]
		if val == target {
			return true
		} else if val < target {
			row++
		} else {
			column--
		}
	}
	return false
}
