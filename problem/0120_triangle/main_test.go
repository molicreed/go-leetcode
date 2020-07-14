package p0120

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect int
	}{
		{
			input: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expect: 11,
		},
		{
			input: [][]int{
				{2},
			},
			expect: 2,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			got := minimumTotal(test.input)
			assert.Equal(t, test.expect, got)
		})
	}
}
