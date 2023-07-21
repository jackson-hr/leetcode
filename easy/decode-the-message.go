package easy

func decodeMessage(key string, message string) string {
	var cur byte = 'a'
	replace := make(map[byte]byte, 27)
	replace[' '] = ' '
	for _, k := range []byte(key) {
		if _, exists := replace[k]; !exists && k != ' ' {
			replace[k] = cur
			cur++
		}
	}

	ans := make([]byte, 0, len(message))
	for _, k := range []byte(message) {
		ans = append(ans, replace[k])
	}
	return string(ans)
}
