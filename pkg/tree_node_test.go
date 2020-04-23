package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMustMakeTreeNode(t *testing.T) {
	var tests = []struct {
		input  string
		expect *TreeNode
	}{
		{
			input: "[1,9,20,null,null,15,7]",
			expect: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			input: "[1,9,20,15,7]",
			expect: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 20,
				},
			},
		},
		{
			input: "[1,2,3, null, null, 4, 5, 6]",
			expect: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			input: "[]",
			expect: nil,
		},
		{
			input: "[1]",
			expect: &TreeNode{
				Val: 1,
			},
		},
	}

	for i, test := range tests {
		got := MustMakeTreeNode(test.input)
		assert.Equal(t, test.expect, got, "case %d", i)
	}
}
