package main

import "fmt"

func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2 := dp[a] * 2
		n3 := dp[b] * 3
		n5 := dp[c] * 5
		dp[i] = n2
		if n3 < dp[i] {
			dp[i] = n3
		}
		if n5 < dp[i] {
			dp[i] = n5
		}
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(nthUglyNumber(n))
}
