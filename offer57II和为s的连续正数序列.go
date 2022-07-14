package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	i, j := 1, 1
	sum := 0
	for i <= target/2 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			arr := make([]int, 0)
			for k := i; k < j; k++ {
				arr = append(arr, k)
			}
			res = append(res, arr)
			sum -= i
			i++
		}
	}
	return res
}

func main() {
	var target int
	fmt.Scan(&target)
	fmt.Println(findContinuousSequence(target))
}
