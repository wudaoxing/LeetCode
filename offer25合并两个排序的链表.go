package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	p := l3
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return l3.Next
}
