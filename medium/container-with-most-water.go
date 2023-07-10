package medium

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans, cur := 0, 0
	for i < j {
		cur = (j - i) * min(height[i], height[j])
		if ans < cur {
			ans = cur
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return ans
}
