package p0002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	resPre := &ListNode{
		Val: 0,
	}
	point := resPre
	for l1 != nil || l2 != nil || carry != 0{
		temp := carry
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		point.Next = &ListNode{
			Val: temp % 10,
		}
		point = point.Next

		carry = temp / 10
	}
	return resPre.Next
}
