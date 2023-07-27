package easy

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for i := range grid {
		sort.Ints(grid[i])
	}

	m := len(grid)
	n := len(grid[0])

	var ans int
	var max int
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		ans += max
		max = 0
	}
	return ans
}
