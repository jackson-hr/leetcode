package easy

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)-1
	if r < 0 || (r == 0 && nums[0] == val) {
		return 0
	}
	for l <= r {
		if nums[l] != val {
			l++
			continue
		}
		if nums[r] == val {
			r--
			continue
		}
		nums[l] = nums[r]
		l++
		r--

	}
	return r + 1
}

func removeElement2(nums []int, val int) int {
	l, r := 0, len(nums)
	if r < 0 {
		return 0
	}
	for l < r {
		if nums[l] != val {
			l++
			continue
		}
		if nums[r-1] == val {
			r--
			continue
		}
		nums[l] = nums[r-1]
		l++
		r--

	}
	return r
}

func removeElement3(nums []int, val int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
