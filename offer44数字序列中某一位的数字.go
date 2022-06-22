package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	//dight:数位；start：该数位初始数字；count：该数位数字个数
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9
	}
	num := start + (n-1)/digit
	str := strconv.Itoa(num)
	return int(str[(n-1)%digit] - '0')
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(findNthDigit(n))
}
