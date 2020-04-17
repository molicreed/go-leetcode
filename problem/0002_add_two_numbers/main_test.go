package p0002

import (
	"testing"
	"strings"
	"strconv"
	// "fmt"

	"github.com/stretchr/testify/assert"
)

type args struct {
	l1 *ListNode
	l2 *ListNode
}

type expect struct {
	val *ListNode
}

func string2listNode(number string) *ListNode {
	var lastNode *ListNode

	nodeValSlice := strings.Split(number, "")

	for _, val := range nodeValSlice {
		nodeVal, _ := strconv.Atoi(val)
		node := &ListNode{
			Val: nodeVal,
			Next: lastNode,
		}
		lastNode = node
	}
	return lastNode
}

func listNode2string(node *ListNode) string {
	res := strconv.Itoa(node.Val)
	for {
		if node = node.Next; node != nil {
			res = strconv.Itoa(node.Val) + res
			continue
		}
		break
	}
	return res
}

func TestAddTowNumbers(t *testing.T) {
	ast := assert.New(t)
	tests := []struct {
		args
		expect
	}{
		{
			args{
				string2listNode("342"), string2listNode("465"),
			}, expect{
				string2listNode("807"),
			},
		},
		{
			args{
				string2listNode("0"), string2listNode("564"),
			}, expect{
				string2listNode("564"),
			},
		},
		{
			args{
				string2listNode("1234"), string2listNode("8766"),
			}, expect{
				string2listNode("10000"),
			},
		},
		{
			args{
				string2listNode("1"), string2listNode("8766"),
			}, expect{
				string2listNode("8767"),
			},
		},
	}

	for _, tt := range tests {
		result := addTwoNumbers(tt.args.l1, tt.args.l2)
		ast.Equal(listNode2string(tt.expect.val), listNode2string(result))
	}
	
}