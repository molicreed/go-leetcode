package leetcode

// import (
// 	"fmt"
// )

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTowNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode
	numbers := []*ListNode{l1, l2}
	point, carry := addTwoNumbersByCarry(l1, l2 , 0)
	result = point
	if isLoopEnd(numbers, carry) {
		return result
	}
	for {
		for i,val:= range numbers{
			if val != nil {
				numbers[i] = val.Next
			}
			
		}

		point.Next, carry = addTwoNumbersByCarry(numbers[0] , numbers[1] , carry)
		

		if isLoopEnd(numbers, carry) {
			break
		}
		point = point.Next
		
	}
	return result
}

func addTwoNumbersByCarry(l1 *ListNode, l2 *ListNode, carry int) (*ListNode, int) {
	var v1,v2 int
	if l1 == nil {
		v1 = 0
	} else {
		v1 = l1.Val
	}
	if l2 == nil {
		v2 = 0
	} else {
		v2 = l2.Val
	}
	value := v1 + v2 + carry

	node := ListNode{value % 10, nil}
	carry = value / 10
	return &node, carry
}

func isLastOne(node *ListNode) bool {
	
	if node == nil {
		return true
	}
	return node.Next == nil
}

func isLoopEnd(nums []*ListNode,carry int) bool {
	bothNil := 0
	for _, val := range nums{
		if isLastOne(val) {
			bothNil++
		}
	}

	return bothNil == 2 && carry == 0
}