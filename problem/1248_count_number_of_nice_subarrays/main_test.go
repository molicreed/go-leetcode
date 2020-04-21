package p1248

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSubarrays(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect int
	}{
		{
			nums:   []int{1, 1, 2, 1, 1},
			k:      3,
			expect: 2,
		},
		{
			nums:   []int{2, 4, 6},
			k:      1,
			expect: 0,
		},
		{
			nums:   []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
			k:      2,
			expect: 16,
		},
	}

	for _, test := range tests {
		got := numberOfSubarrays(test.nums, test.k)
		assert.Equal(t, test.expect, got, "nums: %v, k: %v", test.nums, test.k)
	}
}
