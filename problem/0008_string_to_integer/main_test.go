package p0008

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  "42",
			expect: 42,
		},
		{
			input:  "   -42",
			expect: -42,
		},
		{
			input:  "4193 with words",
			expect: 4193,
		},
		{
			input:  "4193 with words",
			expect: 4193,
		},
		{
			input:  "words and 987",
			expect: 0,
		},
		{
			input:  "-91283472332",
			expect: -2147483648,
		},
		{
			input:  "+1",
			expect: 1,
		},
		{
			input:  "+-2",
			expect: 0,
		},
		{
			input:  "   +0 123",
			expect: 0,
		},
		{
			input:  "123-",
			expect: 123,
		},
		{
			input:  " -1010023630o4",
			expect: -1010023630,
		},
		{
			input:  "2147483648",
			expect: 2147483647,
		},
		{
			input:  "9223372036854775808",
			expect: 2147483647,
		},
		{
			input:  "  0000000000012345678",
			expect: 12345678,
		},
	}

	for _, test := range tests {
		got := myAtoi(test.input)
		assert.Equal(t, test.expect, got)
	}
}
