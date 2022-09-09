package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1       //链表长度
	iter := head //迭代节点
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n -
}
