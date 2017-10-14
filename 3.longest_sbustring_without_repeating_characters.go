package leetcode


func Start3(s string) int {
	return lengthOfLongestSubstring(s)
}

// Given a string, find the length of the longest substring without repeating characters.

// Examples:
// Given "abcabcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var strArr [128]int
	max := 0
	for i, j := 0, 0; j < n; j++ {
		if i < strArr[s[j]]{
			i = strArr[s[j]]
		}
		if max < j-i+1 {
			max = j-i+1
		}
		strArr[s[j]] = j+1
	}
	return max
}
