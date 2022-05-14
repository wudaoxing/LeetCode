package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	queue := make([]TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return nil
	}
	queue = append(queue, *root)

	for len(queue) != 0 {
		temp := queue[0]
		res = append(res, temp.Val)
		queue = queue[1:]
		if temp.Left != nil {
			queue = append(queue, *temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, *temp.Right)
		}
	}
	return res
}

func main() {
	root := TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	fmt.Println(levelOrder(&root))
}
