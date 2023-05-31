package easy

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	m := 0
	for l <= r {
		m = (r-l)>>1 + l
		if m*m <= x {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}

// mySqrt2 sqrt by official
func mySqrt2(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
