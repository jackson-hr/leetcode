package medium

func vowelStrings(words []string, queries [][]int) []int {
	cnt := make([]int, len(words))
	ans := make([]int, len(queries))

	sum := 0
	for i, v := range words {
		sum = 0
		if i > 0 {
			sum += cnt[i-1]
		}
		if flags[v[0]] && flags[v[len(v)-1]] {
			sum++
		}
		cnt[i] = sum
	}

	for i, v := range queries {
		if v[0] > 0 {
			ans[i] = cnt[v[1]] - cnt[v[0]-1]
		} else {
			ans[i] = cnt[v[1]]
		}
	}
	return ans
}

var flags = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}
