package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//拼接+拆分
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	//构建拼接链表
	cur := head
	for cur != nil {
		temp := &Node{cur.Val, nil, nil}
		temp.Next = cur.Next
		cur.Next = temp
		cur = cur.Next.Next
	}
	//构建新节点的random指向
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	//拆分链表
	cur = head.Next
	pre := head
	res := head.Next
	for cur.Next != nil {
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = nil
	return res
}
