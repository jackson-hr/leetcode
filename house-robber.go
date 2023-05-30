package leetcode

func rob(nums []int) int {
	n := len(nums)
	cash := make([]int, n)
	cash[0] = nums[0]
	for i := 1; i < n; i++ {
		if i > 1 && cash[i-2]+nums[i] > cash[i-1] {
			cash[i] = cash[i-2] + nums[i]
			continue
		}

		if cash[i-1] > nums[i] {
			cash[i] = cash[i-1]
		} else {
			cash[i] = nums[i]
		}
	}
	return cash[n-1]
}

func rob2(nums []int) int {
	cash := []int{nums[0], nums[0]}
	for i := 1; i < len(nums); i++ {
		cur := cash[0] + nums[i]
		if i > 1 && cur > cash[1] {
			cash[0] = cash[1]
			cash[1] = cur
			continue
		}
		cash[0] = cash[1]
		if cash[1] < nums[i] {
			cash[1] = nums[i]
		}
	}
	return cash[1]
}
