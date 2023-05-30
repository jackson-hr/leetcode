package easy

var ref = map[byte]byte{
	'}': '{',
	')': '(',
	']': '[',
}

func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	stack := make([]byte, half+1)
	ind := 0
	var rp byte
	for _, p := range []byte(s) {
		if p == '[' || p == '(' || p == '{' {
			stack[ind] = p
			ind++
			if ind > half {
				return false
			}
		} else {
			rp = ref[p]
			ind--
			if ind < 0 || stack[ind] != rp {
				return false
			}
		}
	}
	if ind != 0 {
		return false
	}
	return true
}

func isValid2(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	stack := make([]byte, n)
	ind := 0
	var rp byte
	for _, p := range []byte(s) {
		if p == '[' || p == '(' || p == '{' {
			stack[ind] = p
			ind++
		} else {
			rp = ref[p]
			ind--
			if ind < 0 || stack[ind] != rp {
				return false
			}
		}
	}
	if ind != 0 {
		return false
	}
	return true
}
