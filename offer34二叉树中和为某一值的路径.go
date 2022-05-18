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

var res = [][]int{}
var path []int

func pathSum(root *TreeNode, target int) [][]int {
	recur(root, target)
	return res
}

func recur(root *TreeNode, target int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	target = target - root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		res = append(res, append([]int(nil), path...))
	}
	recur(root.Left, target)
	recur(root.Right, target)
	path = path[:len(path)-1]
}

func creatTree(arr []string, start int) *TreeNode {
	if arr[start] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(arr[start])
	root := &TreeNode{val, nil, nil}
	lnode := 2*start + 1
	rnode := 2*start + 2
	if lnode > len(arr)-1 {
		root.Left = nil
	} else {
		root.Left = creatTree(arr, lnode)
	}
	if rnode > len(arr)-1 {
		root.Right = nil
	} else {
		root.Right = creatTree(arr, rnode)
	}
	return root
}

func main() {
	arr := make([]string, 0)
	var target int
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		arr = strings.Split(sca.Text(), ",")
	}
	root := creatTree(arr, 0)
	fmt.Scan(&target)
	fmt.Println(pathSum(root, target))
}

//奇怪的是力扣上面的不能通过，需要把res和path声明为局部变量才可以，也就是下面这个版本可以通过
//func pathSum(root *TreeNode, target int) [][]int {
//	if root == nil{
//		return nil
//	}
//	var res [][]int
//	var path []int
//	recur(root, target, &res, path)
//	return res
//}
//
//func recur(root *TreeNode, target int, res *[][]int, path []int ) {
//	if root == nil {
//		return
//	}
//	path = append(path, root.Val)
//	target = target - root.Val
//	if target == 0 && root.Left == nil && root.Right == nil {
//		*res = append(*res, append([]int(nil), path...))
//	}
//	recur(root.Left, target, res, path)
//	recur(root.Right, target,res, path)
//	path = path[:len(path)-1]
//}
