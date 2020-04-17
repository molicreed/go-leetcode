package p0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	nums []int
	target int
}

func TestTwoSum(t *testing.T) {
	ast := assert.New(t)
	tests := []struct {
		args    args
		wantResult []int
	}{
		{args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{args{[]int{2, 7, 11, 15}, 13}, []int{0, 2}},
		{args{[]int{2, 7, 11, 15, 19}, 22}, []int{1, 3}},
	}

	for _, tt := range tests {
		result := twoSum(tt.args.nums, tt.args.target)
		ast.Equal(tt.wantResult, result)
	}
	
}