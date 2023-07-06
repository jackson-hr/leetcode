package medium

func numMovesStones(a int, b int, c int) []int {
	x := min(a, min(b, c))
	z := max(a, max(b, c))
	y := a + b + c - x - z

	ans := []int{2, 2}
	if z-y == 1 && y-x == 1 {
		ans = []int{0, 0}
	} else if z-y <= 2 || y-x <= 2 {
		ans = []int{1, 0}
	}
	ans[1] = z - x - 2
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
