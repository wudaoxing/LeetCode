package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归
//func trimBST(root *TreeNode, low int, high int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	if root.Val < low {
//		return  trimBST(root.Right, low, high)
//	} else if root.Val > high {
//		return  trimBST(root.Left, low, high)
//	}
//	root.Left = trimBST(root.Left,low,high)
//	root.Right = trimBST(root.Right,low,high)
//	return root
//}
//迭代
func trimBST(root *TreeNode, low, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	for node := root; node.Left != nil; {
		if node.Left.Val < low {
			node.Left = node.Left.Right
		} else {
			node = node.Left
		}
	}
	for node := root; node.Right != nil; {
		if node.Right.Val > high {
			node.Right = node.Right.Left
		} else {
			node = node.Right
		}
	}
	return root
}
