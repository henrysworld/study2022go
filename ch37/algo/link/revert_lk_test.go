package link

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRevertLk(t *testing.T) {
}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr.Next != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
