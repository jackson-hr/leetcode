package medium

func minFallingPathSum(matrix [][]int) int {
	flags := [2][]int{
		matrix[0],
	}
	n := len(matrix)
	flags[1] = make([]int, n)

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				flags[1][j] = min(flags[0][j], flags[0][j+1]) + matrix[i][j]
			} else if j == n-1 {
				flags[1][j] = min(flags[0][j], flags[0][j-1]) + matrix[i][j]
			} else {
				flags[1][j] = min(min(flags[0][j-1], flags[0][j]), flags[0][j+1]) + matrix[i][j]
			}
		}
		copy(flags[0], flags[1])
	}

	ans := flags[0][0]
	for i := 0; i < n; i++ {
		if flags[0][i] < ans {
			ans = flags[0][i]
		}
	}

	return ans
}

func minFallingPathSum1(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	dp[0] = matrix[0]

	for i := 1; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == n-1 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i-1][j+1]) + matrix[i][j]
			}
		}
	}

	minVal := dp[n-1]
	ans := minVal[0]
	for i := 0; i < n; i++ {
		if minVal[i] < ans {
			ans = minVal[i]
		}
	}

	return ans
}
