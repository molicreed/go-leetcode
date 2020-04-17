package p0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	nums1 []int
	nums2 []int
}

func TestFindMedianSortedArrays(t *testing.T) {
	ast := assert.New(t)
	tests := []struct {
		arg    args
		expect float64
	}{
		{
			args{ 
				[]int{3}, 
				[]int{},
			}, 
			3,
		},
		{
			args{ 
				[]int{}, 
				[]int{3, 4},
			}, 
			3.5,
		},
		{
			args{ 
				[]int{1, 3}, 
				[]int{2},
			}, 
			2.0,
		},
		{
			args{ 
				[]int{1, 2}, 
				[]int{3, 4},
			}, 
			2.5,
		},
		{
			args{ 
				[]int{1, 3, 5, 7, 9}, 
				[]int{2, 4, 6, 8, 10},
			}, 
			5.5,
		},
	}

	for _, tt := range tests {
		result := findMedianSortedArrays(tt.arg.nums1, tt.arg.nums2)
		ast.Equal(tt.expect, result)
	}
	
}