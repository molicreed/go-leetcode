package p0102

import (
	"github.com/stretchr/testify/assert"
	"github.com/yxlimo/leetcode-go/pkg"
	"testing"
)

func TestLevelOrder(t *testing.T) {

	tests := []struct {
		tree   string
		expect [][]int
	}{
		{
			tree: "[1, 9, 20, null, 15, null, 7]",
			expect: [][]int{
				{1},
				{9, 20},
				{15, 7},
			},
		},
	}
	for i, test := range tests {
		root := pkg.MustMakeTreeNode(test.tree)
		got := levelOrder(root)
		assert.Equal(t, test.expect, got, "case %d", i)
	}
}
