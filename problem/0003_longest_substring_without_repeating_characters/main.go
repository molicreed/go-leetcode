package p0003

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	// strArr 记录对应字符开头的最长长度
	var strArr [128]int
	max := 0
	for i, j := 0, 0; j < n; j++ {
		if i < strArr[s[j]]{
			i = strArr[s[j]]
		}
		// j-i+1 为当前统计到的不重复子串的长度
		if max < j-i+1 {
			max = j-i+1
		}
		strArr[s[j]] = j+1
	}
	return max
}
