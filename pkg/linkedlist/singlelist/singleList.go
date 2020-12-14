package singlelist

import "strconv"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode ..
type ListNode struct {
	Val  int
	Next *ListNode
}

// Add String method to use print
func (l *ListNode) String() string {
	var res string
	for l != nil {
		res += strconv.Itoa(l.Val)
		res += "->"
		l = l.Next
	}
	res += "nil"
	return res
}

//AddSingleList ..
func (list *ListNode) AddSingleList(l *ListNode, Val int) *ListNode {
	if l.Next == nil && l.Val == 0 {
		l.Val = Val
		return l
	}

	temp := new(ListNode)
	temp.Val = Val
	temp.Next = nil

	for l.Next != nil {
		l = l.Next
	}
	l.Next = temp
	return l
}
