package p0171

import (
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n-1][m], dp[n][m-1] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
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
