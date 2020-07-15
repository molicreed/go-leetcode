package p0096

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{
			input:  0,
			expect: 1,
		},
		{
			input:  1,
			expect: 1,
		},
		{
			input:  2,
			expect: 2,
		},
		{
			input:  3,
			expect: 5,
		},
	}

	for _, test := range tests {
		got := numTrees(test.input)
		assert.Equal(t, test.expect, got, fmt.Sprintf("input: %v", test.input))
	}
}
