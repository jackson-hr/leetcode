package medium

func maxAbsoluteSum(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}

	preSum := make([]int, length)
	preSum[0] = nums[0]
	for i := 1; i < length; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	max, min := 0, 0
	for _, s := range preSum {
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}

	return max - min
}
