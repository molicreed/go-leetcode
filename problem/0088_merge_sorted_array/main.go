package p0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, p := m-1, n-1, m+n-1
	for i > -1 || j > -1 {
		if i > -1 && (j == -1 || nums1[i] >= nums2[j]) {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}
}
