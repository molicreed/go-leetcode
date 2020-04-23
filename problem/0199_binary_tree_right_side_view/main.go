package p0199

import (
	"container/list"
	. "github.com/yxlimo/leetcode-go/pkg"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	l := list.New()
	result := make([]int, 0, 20)
	l.PushBack(root)
	temp := make([]*TreeNode, 0, 100)
	for l.Len() != 0 {
		for l.Front() != nil {
			e := l.Front()
			l.Remove(e)
			node := e.Value.(*TreeNode)
			temp = append(temp, node)
		}
		if len(temp) > 0 {
			result = append(result, temp[len(temp)-1].Val)
		}
		for _, node := range temp {
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		temp = temp[0:0]
	}
	return result
}
