package p0171

func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n)
	for i := 0; i <= n-1; i++ {
		dp[i] = make([]int, m)
	}
	dp[n-1][m-1] = health(1 - dungeon[n-1][m-1])
	// 最后一列
	for i, j := n-2, m-1; i >= 0; i-- {
		dp[i][j] = health(dp[i+1][j] - dungeon[i][j])
	}
	// 最后一行
	for i, j := n-1, m-2; j >= 0; j-- {
		dp[i][j] = health(dp[i][j+1] - dungeon[i][j])
	}
	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			dp[i][j] = health(min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j])
		}
	}
	return dp[0][0]
}

func health(result int) int {
	if result >= 1 {
		return result
	}
	return 1
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
