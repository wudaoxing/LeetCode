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
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	count := 0
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
		if count%2 == 0 {
			res = append(res, temp)
		} else {
			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp2 := temp[i]
				temp[i] = temp[j]
				temp[j] = temp2
			}
			res = append(res, temp)
		}
		count++
	}
	return res
}

func createTree(arr []string, start int) *TreeNode {
	if arr[start] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(arr[start])
	root := &TreeNode{val, nil, nil}
	lnode := start*2 + 1
	rnode := start*2 + 2
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
