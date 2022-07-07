package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := depth(root)
	if res == -1 {
		return false
	}
	return true
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	if left == -1 {
		return -1
	}
	right := depth(root.Right)
	if right == -1 {
		return -1
	}
	if math.Abs(float64(left-right)) > 1 {
		return -1
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

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
	fmt.Println(isBalanced(root))
}
