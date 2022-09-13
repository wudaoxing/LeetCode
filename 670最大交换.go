package main

import (
	"fmt"
	"strconv"
)

//func maximumSwap(num int) int {
//	if num < 12 {
//		return num
//	}
//	tmp := num
//	arr := make([]int, 0)
//	for i := 0; tmp > 0; i++ {
//		arr = append(arr, tmp%10)
//		tmp /= 10
//	}
//	idx := make([]int, len(arr))
//	for i := 0; i < len(arr); i++ {
//		idx[i] = i
//	}
//	sort.Slice(idx, func(i, j int) bool {
//		a, b := idx[i], idx[j]
//		return arr[a] > arr[b]
//	})
//	for i := 0; i < len(arr); i++ {
//		for arr[idx[i]] == arr[len(arr)-1-i] {
//			arr[idx[i]], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[idx[i]]
//			break
//		}
//	}
//	ans := 0
//	for i := len(arr) - 1; i >= 0; i-- {
//		ans += arr[i] * pow(i)
//	}
//	return ans
//}
//
//func pow(i int) int {
//	ans := 1
//	for i > 0 {
//		ans *= 10
//		i--
//	}
//	return ans
//}

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	id1 := -1
	id2 := -1
	maxID := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > s[maxID] {
			maxID = i
		} else if s[i] < s[maxID] {
			id1 = i
			id2 = maxID
		}
	}
	if id1 < 0 {
		return num
	}
	s[id1], s[id2] = s[id2], s[id1]
	ans, _ := strconv.Atoi(string(s))
	return ans
}
func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(maximumSwap(num))
}
