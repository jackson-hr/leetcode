package medium

import "fmt"

func equalPairs(grid [][]int) int {
	ans := 0
	n := len(grid)
	rows := make(map[string]int, n)
	for _, v := range grid {
		rows[fmt.Sprint(v)]++
	}
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[j] = grid[j][i]
		}
		if subAns, exists := rows[fmt.Sprint(tmp)]; exists {
			ans += subAns
		}
	}
	return ans
}
