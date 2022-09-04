package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//func connect(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//	st := make([]*Node, 0)
//	st = append(st, root)
//	nilNode := &Node{Val: 1001}
//	st = append(st, nilNode)
//	i := 1
//	for len(st) != 0 {
//		node := st[0]
//		if len(st) > 1 && st[1].Val == 1001 {
//			node.Next = nil
//			st = st[2:]
//		} else {
//			st = st[1:]
//			node.Next = st[0]
//		}
//		if node.Left != nil {
//			st = append(st, node.Left)
//			st = append(st, node.Right)
//			i += 2
//			if judge(i) {
//				st = append(st, nilNode)
//			}
//		}
//	}
//	return root
//}
//
//func judge(length int) bool {
//	for i := 1; i <= 12; i++ {
//		if length == int(math.Pow(2.0, float64(i)))-1 {
//			return true
//		}
//	}
//	return false
//}

//使用已建立的next指针
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	for leftMost := root; leftMost.Left != nil; leftMost = leftMost.Left {
		for node := leftMost; node != nil; node = node.Next {
			node.Left.Next = node.Right

			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

func buildTree(a []string, index int) *Node {
	if a[index] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(a[index])
	root := &Node{val, nil, nil, nil}
	lnode := 2*index + 1
	rnode := 2*index + 2
	if lnode >= len(a) {
		root.Left = nil
	} else {
		root.Left = buildTree(a, lnode)
	}
	if rnode >= len(a) {
		root.Right = nil
	} else {
		root.Right = buildTree(a, rnode)
	}
	return root
}

func main() {
	var a []string
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		a = strings.Split(sca.Text(), ",")
	}
	root := buildTree(a, 0)
	fmt.Println(connect(root))
}
