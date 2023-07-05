package easy

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	var ans = numOnes
	k -= numOnes
	if k <= 0 {
		ans += k
	}
	k -= numZeros
	if k <= 0 {
		return ans
	}
	ans += -1 * k
	return ans
}
