package hard

func trap(height []int) int {
	n := len(height)
	lm, rm := height[0], height[n-1]
	left := make([]int, n)
	right := make([]int, n)

	for i := 1; i < n; i++ {
		if height[i] < lm {
			left[i] = lm - height[i]
		} else {
			lm = height[i]
		}
	}
	for j := n - 2; j >= 0; j-- {
		if height[j] < rm {
			right[j] = rm - height[j]
		} else {
			rm = height[j]
		}
	}

	var ans = 0
	for i := 0; i < n; i++ {
		ans += min(left[i], right[i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
