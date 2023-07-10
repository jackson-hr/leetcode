package medium

func twoSum(numbers []int, target int) []int {
	ans := []int{0, len(numbers) - 1}
	var res = 0
	for ans[0] < ans[1] {
		res = numbers[ans[0]] + numbers[ans[1]]
		if res == target {
			break
		}

		if res > target {
			ans[1]--
		} else {
			ans[0]++
		}
	}
	ans[0]++
	ans[1]++
	return ans
}
