package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	//if val > root.Val {
	//	return &TreeNode{
	//		val, root, nil,
	//	}
	//}
	//pre := root
	//cur := root
	//for cur != nil && cur.Val > val {
	//	pre = cur
	//	cur = cur.Right
	//}
	//pre.Right = &TreeNode{val, cur, nil}
	//return root
	var parent *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{val, root, nil}
			}
			parent.Right = &TreeNode{val, cur, nil}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{Val: val}
	return root
}

func buildTree(a []string, index int) *TreeNode {
	if a[index] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(a[index])
	root := &TreeNode{val, nil, nil}
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
func levelTraverse(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	a := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		if tmp.Val == 0 {
			a = append(a, "null")
			continue
		}
		val := strconv.Itoa(tmp.Val)
		a = append(a, val)
		if tmp.Left == nil && tmp.Right == nil {
			continue
		}
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		} else {
			queue = append(queue, &TreeNode{Val: 0})
		}
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		} else {
			queue = append(queue, &TreeNode{Val: 0})
		}
	}
	if a[len(a)-1] == "null" {
		a = a[:len(a)-1]
	}
	return a
}
func main() {
	var val int
	var a []string
	fmt.Print("val=")
	fmt.Scanf("%d\n", &val)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		a = strings.Split(sca.Text(), ",")
	}
	root := buildTree(a, 0)
	res := insertIntoMaxTree(root, val)
	fmt.Println(levelTraverse(res))
}
