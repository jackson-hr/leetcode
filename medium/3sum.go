package medium

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	res := make([][]int, 0, 4)
	for l := 0; l < n-2; l++ {
		// skip repeat
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}

		if nums[l]+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		target := -1 * nums[l]
		r := n - 1

		for m := l + 1; m < r; m++ {
			// skip repeat
			if m > l+1 && nums[m] == nums[m-1] {
				continue
			}

			if nums[l]+nums[m]+nums[m+1] > 0 {
				break
			}

			// reduce range
			for m < r && nums[m]+nums[r] > target {
				r--
			}

			if m == r {
				break
			}

			if nums[m]+nums[r] == target {
				res = append(res, []int{nums[l], nums[m], nums[r]})
			}
		}
	}
	return res
}
