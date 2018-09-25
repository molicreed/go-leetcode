package p0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	ast := assert.New(t)
	tests := []struct {
		arg    string
		expect int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"enciodfvndjasf", 8},
	}

	for _, tt := range tests {
		result := lengthOfLongestSubstring(tt.arg)
		ast.Equal(tt.expect, result)
	}
	
}