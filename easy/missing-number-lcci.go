package easy

func missingNumber(nums []int) int {
	n := len(nums)
	sum := (n + 1) * n / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
