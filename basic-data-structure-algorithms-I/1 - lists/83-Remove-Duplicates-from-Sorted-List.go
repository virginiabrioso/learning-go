package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}

	deleteDuplicates(head)
	printList(head)
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	current := head

	// processing each node so O(n) time complexity
	// no additional data so O(1) space complexity
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head

}
