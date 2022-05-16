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

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		temp := make([]int, 0)
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func createTree(arr []string, start int) *TreeNode {
	if arr[start] == "null" {
		return nil
	}
	//新建根结点
	val, _ := strconv.Atoi(arr[start])
	root := &TreeNode{val, nil, nil}
	//左右节点索引
	lnode := 2*start + 1
	rnode := 2*start + 2

	if lnode > len(arr)-1 {
		root.Left = nil
	} else {
		root.Left = createTree(arr, lnode)
	}

	if rnode > len(arr)-1 {
		root.Right = nil
	} else {
		root.Right = createTree(arr, rnode)
	}

	return root
}

func main() {
	arr := make([]string, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		arr = strings.Split(sca.Text(), ",")
	}
	root := createTree(arr, 0)
	fmt.Println(levelOrder(root))
}
