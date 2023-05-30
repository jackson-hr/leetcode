package easy

func averageValue(nums []int) int {
	sum := 0
	cnt := 0
	for _, v := range nums {
		if v%6 == 0 {
			sum += v
			cnt++
		}
	}
	if cnt < 1 {
		return 0
	}

	return sum / cnt
}
