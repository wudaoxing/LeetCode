package main

import "fmt"

func lastRemaining(n int, m int) int {
	tmp := 0
	for i := 2; i <= n; i++ {
		tmp = (tmp + m) % i
	}
	return tmp
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Println(lastRemaining(n, m))
}
