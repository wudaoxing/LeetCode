package main

import "fmt"

func sumNums(n int) int {
	res := 0
	return sum(n, &res)
}

func sum(n int, res *int) int {
	x := n > 1 && sum(n-1, res) > 0
	x = !x
	*res += n
	return *res
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sumNums(n))
}
