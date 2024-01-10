package easy

func minLength(s string) int {
	var stack []byte
	for i := range s {
		stack = append(stack, s[i])
		leng := len(stack)
		if leng >= 2 && (string(stack[leng-2:]) == "AB" || string(stack[leng-2:]) == "CD") {
			stack = stack[:leng-2]
		}
	}
	return len(stack)
}
