package leetcode


func TwoSum(nums []int , target int) []int {
	numsMap := make(map[int]int, len(nums))

	for i, val := range nums {
		complement := target - val
		if key, ok := numsMap[complement]; ok {
			return  []int{i, key}
		}
		numsMap[val] = i
	}
	return nil
}

