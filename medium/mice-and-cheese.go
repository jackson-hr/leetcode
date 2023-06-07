package medium

import (
	"sort"
)

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	diffs := make([]int, n)
	var ans = 0
	for i := 0; i < n; i++ {
		diffs[i] = reward1[i] - reward2[i]
		ans += reward2[i]
	}
	sort.Ints(diffs)
	for i := k; i > 0; i-- {
		ans += diffs[n-i]
	}
	return ans
}
