package p0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	totalLength := len1 + len2
	var p1, p2 int
	len0 := totalLength / 2 + 1

	var nowData, preData int
	getP1 := func() {
		nowData =  nums1[p1]
		p1++
	}
	getP2 := func() {
		nowData =  nums2[p2]
		p2++
	}
	for i := 0; i < len0; i++ {
		preData = nowData
		if p1 >= len1 {
			getP2()
			continue
		}
		if p2 >= len2 {
			getP1()
			continue
		}

		if nums1[p1] >= nums2[p2] {
			getP2()
		} else {
			getP1()
		}
	}
	
	if totalLength % 2 != 0 {
		return float64(nowData)
	}
	return float64(nowData + preData) / 2
}