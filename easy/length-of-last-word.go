package easy

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if char == ' ' && length != 0 {
			return length
		}
		if char != ' ' {
			length++
		}
	}
	return length
}
