package easy

func applyOperations(nums []int) []int {
	n := len(nums)
	emptyIdx := 0
	for i := 0; i < n; i++ {
		if i < n-1 && nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
		if nums[i] != 0 {
			nums[i], nums[emptyIdx] = nums[emptyIdx], nums[i]
			emptyIdx++
		}
	}
	return nums
}
