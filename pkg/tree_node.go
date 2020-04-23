package pkg

import (
	"encoding/json"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MakeTreeNode accept array like "[1,2,3,null,5,null,4]"
func MustMakeTreeNode(treeStr string) *TreeNode {
	treeSlice := make([]int, 0)
	if err := json.Unmarshal([]byte(treeStr), &treeSlice); err != nil {
		panic("illegal input")
	}
	if len(treeSlice) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: treeSlice[0],
	}
	temp := append(make([]*TreeNode, 0, len(treeSlice)), root)
	nextLevel := make([]*TreeNode, 0, 1)
	i := 0
	for len(temp) != 0{
		for _, v := range temp {

			if i++; i < len(treeSlice) && treeSlice[i] != 0 {
				v.Left = &TreeNode{
					Val: treeSlice[i],
				}
				nextLevel = append(nextLevel, v.Left)
			}

			if i++; i < len(treeSlice) && treeSlice[i] != 0 {
				v.Right = &TreeNode{
					Val: treeSlice[i],
				}
				nextLevel = append(nextLevel, v.Right)
			}
		}
		temp = nextLevel
		nextLevel = make([]*TreeNode, 0, len(temp) * 2)
	}

	return root
}

