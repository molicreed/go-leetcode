package p1248

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	oddIndex := make([]int, 0, len(nums))
	for i, v := range nums {
		if in&0x1 != 0 {
			oddIndex = append(oddIndex, i)
		}
	}
	if len(oddIndex) < k {
		return 0
	}
	result := 0
	for first := 0; first < len(oddIndex)-k+1; first++ {
		countPrefix := 0
		if first == 0 {
			countPrefix = oddIndex[first] + 1
		} else {
			countPrefix = oddIndex[first] - oddIndex[first-1]
		}
		last := first + k - 1
		countSuffix := 0
		if last == len(oddIndex)-1 {
			countSuffix = len(nums) - oddIndex[last]
		} else {
			countSuffix = oddIndex[last+1] - oddIndex[last]
		}
		result += countPrefix * countSuffix
	}
	return result
}