package medium

func maximumEvenSplit(finalSum int64) []int64 {
	ans := make([]int64, 0)
	if finalSum%2 != 0 {
		return ans
	}

	for i := int64(2); finalSum > 0; i += 2 {
		finalSum -= i
		if finalSum <= i {
			i += finalSum
			finalSum = 0
		}
		ans = append(ans, i)
	}
	return ans
}

func maximumEvenSplit1(finalSum int64) []int64 {
	ans := make([]int64, 0)
	if finalSum%2 != 0 {
		return ans
	}

	for i := int64(2); finalSum >= i; i += 2 {
		finalSum -= i
		ans = append(ans, i)
	}
	ans[len(ans)-1] += finalSum
	return ans
}
