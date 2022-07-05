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

var res, K int

func kthLargest(root *TreeNode, k int) int {
	K = k
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	if K == 0 {
		return
	}
	K--
	if K == 0 {
		res = root.Val
	}
	dfs(root.Left)
}

func buildTree(arr []string, index int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	if arr[index] == "null" {
		return nil
	}
	v, _ := strconv.Atoi(arr[index])
	root := &TreeNode{v, nil, nil}
	//左节点索引
	lnode := 2*index + 1
	if lnode < len(arr) {
		root.Left = buildTree(arr, lnode)
	}
	//右节点索引
	rnode := 2*index + 2
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
	var k int
	if sca.Scan() {
		k, _ = strconv.Atoi(sca.Text())
	}
	root := buildTree(arr, 0)
	fmt.Println(kthLargest(root, k))
}
