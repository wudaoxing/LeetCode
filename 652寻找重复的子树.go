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

// 序列化子树为字符串，map字符串与*TreeNode
//func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
//	repeat := map[*TreeNode]struct{}{}
//	seen := map[string]*TreeNode{}
//	var dfs func(*TreeNode) string
//	dfs = func(node *TreeNode) string {
//		if node == nil {
//			return ""
//		}
//		serial := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
//		if n, ok := seen[serial]; ok {
//			repeat[n] = struct{}{}
//		} else {
//			seen[serial] = node
//		}
//		return serial
//	}
//	dfs(root)
//	ans := make([]*TreeNode, 0, len(repeat))
//	for node := range repeat {
//		ans = append(ans, node)
//	}
//	return ans
//}

//三元组（x，i，j）：x表示该子树的根结点的值，i表示左子树的序号，j表示右子树的序号。每个子树的序号表示该子树结构，结构同序号同，三元组一致表示结构相同。
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	type pair struct {
		node *TreeNode
		idx  int
	}
	repeat := map[*TreeNode]struct{}{}
	seen := map[[3]int]pair{}
	idx := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
		if p, ok := seen[tri]; ok {
			repeat[p.node] = struct{}{}
			return p.idx
		}
		idx++
		seen[tri] = pair{node, idx}
		return idx
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
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
	var a []string
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		a = strings.Split(sca.Text(), ",")
	}
	root := buildTree(a, 0)
	fmt.Println(findDuplicateSubtrees(root))
}
