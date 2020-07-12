package p0171

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateMinimumHP(t *testing.T) {

	tests := []struct {
		dungeon [][]int
		expect  int
	}{
		{
			dungeon: [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
			expect: 7,
		},
		{
			dungeon: [][]int{
				{100},
			},
			expect: 1,
		},
	}
	for i, test := range tests {
		got := calculateMinimumHP(test.dungeon)
		assert.Equal(t, test.expect, got, "case %d", i)
	}
}
