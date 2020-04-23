package interview_08_11

const (
	mode = 1000000007

	quarters = 25
	dimes    = 10
	nickels  = 5
)

func waysToChange_1(n int) int {
	quartersMax := n / 25
	ans := 0
	for i := 0; i <= quartersMax; i++ {
		remain := n - i*quarters
		dimesMax := remain / dimes
		nickelsMax := (remain - dimesMax*dimes) / nickels
		ans += (dimesMax + 1) * (dimesMax + nickelsMax + 1)
	}
	return ans % mode
}

// TODO 动态规划
func waysToChange_2(n int) int {
	return 0
}
