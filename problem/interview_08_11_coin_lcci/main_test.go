package interview_08_11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaysToChange(t *testing.T) {
	var tests = []struct {
		n      int
		expect int
	}{
		{
			n:      5,
			expect: 2,
		},
		{
			n:      10,
			expect: 4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, waysToChange_1(test.n))
	}
}
