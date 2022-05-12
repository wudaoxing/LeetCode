package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (isMatch(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func isMatch(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}

	return isMatch(A.Left, B.Left) && isMatch(A.Right, B.Right)
}
