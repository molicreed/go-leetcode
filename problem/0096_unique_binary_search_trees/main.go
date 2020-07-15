package p0096

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		for left := 0; left < i; left++ {
			f[i] += f[left] * f[i-left-1]
		}
	}
	return f[n]
}
