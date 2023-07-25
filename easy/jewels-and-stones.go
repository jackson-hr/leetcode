package easy

func numJewelsInStones(jewels string, stones string) int {
	flags := make(map[rune]bool, len(jewels))
	for _, s := range jewels {
		flags[s] = true
	}

	ans := 0
	for _, s := range stones {
		if _, exists := flags[s]; exists {
			ans++
		}
	}
	return ans
}
