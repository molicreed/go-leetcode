package p0120

func minimumTotal(triangle [][]int) int {
	p := make([]int, len(triangle))
	p[0] = triangle[0][0]
	offset := 1
	for offset < len(p) {
		p[offset] = p[offset-1] + triangle[offset][offset]
		for i := offset - 1; i >= 1; i-- {
			p[i] = min(p[i], p[i-1]) + triangle[offset][i]
		}
		p[0] = p[0] + triangle[offset][0]
		offset++
	}
	result := p[0]
	for i := 1; i < len(p); i++ {
		result = min(result, p[i])
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
