package p0199

import (
	"github.com/stretchr/testify/assert"
	"github.com/yxlimo/leetcode-go/pkg"
	"testing"
)

func TestRightSideView(t *testing.T) {

	tests := []struct {
		tree   string
		expect []int
	}{
		{
			tree:   "[1,2,3,null,5,null,4]",
			expect: []int{1, 3, 4},
		},
		{
			tree:   "[1, 2, 3, 4]",
			expect: []int{1, 3, 4},
		},
		{
			tree:   "[1,2,3, null, null, 4, 5, 6]",
			expect: []int{1, 3, 5, 6},
		},
	}

	for _, test := range tests {
		treeNode := pkg.MustMakeTreeNode(test.tree)

		got := rightSideView(treeNode)
		assert.Equal(t, test.expect, got, "tree: %v", test.tree)
	}
}
