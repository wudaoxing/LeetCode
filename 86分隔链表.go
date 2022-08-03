package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//官方答案
//func partition(head *ListNode, x int) *ListNode {
//	small := &ListNode{}
//	smallHead := small
//	large := &ListNode{}
//	largeHead := large
//	for head != nil {
//		if head.Val < x {
//			small.Next = head
//			small = small.Next
//		} else {
//			large.Next = head
//			large = large.Next
//		}
//		head = head.Next
//	}
//	large.Next = nil
//	small.Next = largeHead.Next
//	return smallHead.Next
//}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	max := head
	pre_ma := head
	for max != nil {
		if max.Val < x {
			pre_ma = max
			max = max.Next
		} else {
			break
		}
		if max.Next == nil {
			return head
		}
	}
	pre_mi := max
	min := max.Next
	for min != nil {
		if min.Val >= x {
			pre_mi = min
			min = min.Next
		} else {
			pre_mi.Next = min.Next
			min.Next = max
			if max == head {
				head = min
			} else {
				pre_ma.Next = min
			}
			pre_ma = min
			min = pre_mi.Next
		}
	}
	return head
}

func buildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{arr[0], nil}
	p := head
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{arr[i], nil}
		p = p.Next
	}
	return head
}
func listToArr(head *ListNode) []int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}

func main() {
	arr := make([]int, 0)
	fmt.Print("head=")
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			arr = append(arr, n)
		}
	}
	fmt.Print("x=")
	var x int
	fmt.Scan(&x)
	head := buildList(arr)
	res := listToArr(partition(head, x))
	fmt.Println(res)
}
