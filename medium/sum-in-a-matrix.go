package medium

import "sort"

func matrixSum(nums [][]int) int {
	first := len(nums)
	second := len(nums[0])

	for i := 0; i < first; i++ {
		sort.Ints(nums[i])
	}
	var ans int
	for j := 0; j < second; j++ {
		max := 0
		for i := 0; i < first; i++ {
			if nums[i][j] > max {
				max = nums[i][j]
			}
		}
		ans += max
	}
	return ans
}
