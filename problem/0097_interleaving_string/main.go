package p0097

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	f := make([][]bool, len(s1)+1)
	for i := range f {
		f[i] = make([]bool, len(s2)+1)
	}
	f[0][0] = true
	for i := 1; i <= len(s1); i++ {
		if f[i-1][0] && s1[i-1] == s3[i-1] {
			f[i][0] = true
		}
	}
	for j := 1; j <= len(s2); j++ {
		if f[0][j-1] && s2[j-1] == s3[j-1] {
			f[0][j] = true
		}
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if f[i-1][j] && s1[i-1] == s3[i+j-1] {
				f[i][j] = true
			} else {
				if f[i][j-1] && s2[j-1] == s3[i+j-1] {
					f[i][j] = true
				}
			}
		}
	}
	return f[len(s1)][len(s2)]
}
