package p0055

func canJump(nums []int) bool {
	index := 0
	maxIndex := 0
	endPoint := len(nums) - 1
	for index <= endPoint {
		if maxIndex < index+nums[index] {
			maxIndex = index + nums[index]
		}
		switch {
		case maxIndex >= endPoint:
			// jump success
			return true
		case maxIndex <= index:
			// can not jump more far
			return false
		default:
			index++
		}
	}
	return false
}
