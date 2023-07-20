package easy

func addStrings(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)

	n := n1
	if n < n2 {
		n = n2
	}
	var tmp = 0
	var carry = 0
	ans := make([]byte, n)
	for i := 1; i <= n; i++ {
		tmp = carry
		if i <= n1 {
			tmp += int(num1[n1-i] - '0')
		}
		if i <= n2 {
			tmp += int(num2[n2-i] - '0')
		}
		carry = tmp / 10
		tmp = tmp % 10
		ans[n-i] = uint8(tmp) + '0'
	}
	if carry > 0 {
		return "1" + string(ans)
	}
	return string(ans)
}
