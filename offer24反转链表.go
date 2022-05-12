package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
双指针法
*/
//func reverseList(head *ListNode) *ListNode {
//	var pre *ListNode = nil
//	cur := head
//
//	for cur != nil {
//		temp := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = temp
//	}
//	return pre
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return temp
}
func main() {
	head := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, &ListNode{
					4, &ListNode{
						5, nil,
					},
				},
			},
		},
	}
	fmt.Println(reverseList(head).Val)
}
