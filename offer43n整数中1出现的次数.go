package main

import "fmt"

func countDigitOne(n int) int {
	res := 0
	high, cur, digit, low := n/10, n%10, 1, 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(countDigitOne(n))
}
