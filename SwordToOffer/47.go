package SwordToOffer

import "fmt"

var max int

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	visit(grid, 0, 0, 0)
	return max
}

func visit(grid [][]int, x, y, res int) {
	if x >= len(grid) || y >= len(grid[0]) {
		if max <= res {
			max = res
		}
		return
	}
	res += grid[x][y]
	visit(grid, x+1, y, res)
	visit(grid, x, y+1, res)
}
