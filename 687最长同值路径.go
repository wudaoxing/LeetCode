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

func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		left1, right1 := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			left1 = left + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right1 = right + 1
		}
		ans = max(ans, left1+right1)
		return max(left1, right1)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
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

func main() {
	strs := make([]string, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs = strings.Split(sca.Text(), ",")
	}
	root := buildTree(strs, 0)
	fmt.Println(longestUnivaluePath(root))
}
