package easy

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func searchInsert2(nums []int, target int) int {
	n := len(nums)
	l, m, r := 0, 0, n-1
	ans := n
	for l <= r {
		m = (r-l)/2 + l
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			ans = m
			r = m - 1
		}
	}
	return ans
}
