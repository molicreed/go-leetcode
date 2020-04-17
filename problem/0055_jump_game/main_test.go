package p0055

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums    []int
		expect bool
	}{
		{
			nums: []int{2,3,1,1,4},
			expect: true,
		},
		{
			nums: []int{3,2,1,0,4},
			expect: false,
		},
		{
			nums: []int{0},
			expect: true,
		},
	}

	for _, test := range tests {
		got := canJump(test.nums)
		assert.Equal(t, test.expect, got, fmt.Sprintf("input: %v", test.nums))
	}
}