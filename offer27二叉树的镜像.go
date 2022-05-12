package main

import "fmt"

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	if root.Left != nil {
		mirrorTree(root.Left)
	}
	if root.Right != nil {
		mirrorTree(root.Right)
	}

	return root
}

func main() {
	var n int //节点个数
	fmt.Scan(&n)
}
