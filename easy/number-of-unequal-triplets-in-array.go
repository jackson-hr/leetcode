package easy

import (
	"sort"
)

func unequalTriplets(nums []int) int {
	sort.Ints(nums)
	res, n := 0, len(nums)
	for i, j := 0, 0; i < n; i = j {
		for j < n && nums[j] == nums[i] {
			j++
		}
		res += i * (j - i) * (n - j)
	}
	return res
}

func unequalTriplets1(nums []int) int {
	cnts := make(map[int]int, 0)
	for _, v := range nums {
		cnts[v]++
	}
	ans, left, n := 0, 0, len(nums)
	for _, cnt := range cnts {
		ans += left * cnt * (n - left - cnt)
		left += cnt
	}
	return ans
}
