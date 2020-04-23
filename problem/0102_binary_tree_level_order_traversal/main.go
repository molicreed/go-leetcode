package p0102

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	l := list.New()
	out := make([][]int, 0, 20)
	l.PushBack(root)
	temp := make([]*TreeNode, 0, 100)
	for l.Len() != 0 {
		out = append(out, []int{})
		for l.Front() != nil {
			e := l.Front()
			l.Remove(e)
			node := e.Value.(*TreeNode)
			temp = append(temp, node)
		}
		for _, node := range temp {
			out[len(out)-1] = append(out[len(out)-1], node.Val)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		temp = temp[0:0]
	}
	return out
}
