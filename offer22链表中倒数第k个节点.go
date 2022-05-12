package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	former := head
	latter := head
	for i := 0; i < k; i++ {
		former = former.Next
	}
	for former != nil {
		latter = latter.Next
		former = former.Next
	}
	return latter
}

func main() {
	head := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(getKthFromEnd(&head, 2).Val)
}
