package p0053

func maxSubArray(nums []int) int {
	pre := nums[0]
	max := pre
	for i := 1; i < len(nums); i++ {
		if pre >= 0 {
			pre = pre + nums[i]
		} else {
			pre = nums[i]
		}
		if pre > max {
			max = pre
		}
	}
	return max
}
