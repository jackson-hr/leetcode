package easy

import "fmt"

func summaryRanges(nums []int) []string {
	n := len(nums)
	ans := make([]string, 0)
	for i := 0; i < n; {
		l := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := fmt.Sprint(nums[l])
		if l < i-1 {
			s += fmt.Sprintf("->%d", nums[i-1])
		}
		ans = append(ans, s)
	}
	return ans
}
