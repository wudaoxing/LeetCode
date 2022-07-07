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

//BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]TreeNode, 0)
	queue = append(queue, *root)
	res := 0
	for len(queue) != 0 {
		res++
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, *node.Left)
			}
			if node.Right != nil {
				queue = append(queue, *node.Right)
			}
		}
	}
	return res
}

//DFS
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	tmp1 := 1
//	if root.Left != nil {
//		tmp1 += maxDepth(root.Left)
//	}
//	tmp2 := 1
//	if root.Right != nil {
//		tmp2 += maxDepth(root.Right)
//	}
//	if tmp1 > tmp2 {
//		return tmp1
//	} else {
//		return tmp2
//	}
//}
func buildTree(arr []string, index int) *TreeNode {
	if arr[index] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(arr[index])
	root := &TreeNode{val, nil, nil}
	lnode := 2*index + 1
	rnode := 2*index + 2
	if lnode < len(arr) {
		root.Left = buildTree(arr, lnode)
	}
	if rnode < len(arr) {
		root.Right = buildTree(arr, rnode)
	}
	return root
}

func main() {
	arr := make([]string, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		arr = strings.Split(sca.Text(), ",")
	}
	root := buildTree(arr, 0)
	fmt.Println(maxDepth(root))
}
