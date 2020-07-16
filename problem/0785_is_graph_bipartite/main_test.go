package p0785

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect bool
	}{
		{
			input: [][]int{
				{1, 3}, {0, 2}, {1, 3}, {0, 2},
			},
			expect: true,
		},
		{
			input: [][]int{
				{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2},
			},
			expect: false,
		},
		{
			input:  [][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}},
			expect: false,
		},
		{
			input:  [][]int{{}, {3}, {}, {1}, {}},
			expect: true,
		},
		{
			input:  [][]int{{2, 3, 5, 6, 7, 8, 9}, {2, 3, 4, 5, 6, 7, 8, 9}, {0, 1, 3, 4, 5, 6, 7, 8, 9}, {0, 1, 2, 4, 5, 6, 7, 8, 9}, {1, 2, 3, 6, 9}, {0, 1, 2, 3, 7, 8, 9}, {0, 1, 2, 3, 4, 7, 8, 9}, {0, 1, 2, 3, 5, 6, 8, 9}, {0, 1, 2, 3, 5, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}},
			expect: false,
		},
	}

	for _, test := range tests {
		got := isBipartite(test.input)
		assert.Equal(t, test.expect, got, fmt.Sprintf("input: %v", test.input))
	}
}
