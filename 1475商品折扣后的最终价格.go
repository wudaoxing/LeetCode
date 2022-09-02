package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func finalPrices(prices []int) []int {
//	for i := 0; i < len(prices)-1; i++ {
//		for j := i + 1; j < len(prices); j++ {
//			if prices[j] <= prices[i] {
//				prices[i] -= prices[j]
//				break
//			}
//		}
//	}
//	return prices
//}

// 单调栈
func finalPrices(prices []int) []int {
	length := len(prices)
	ans := make([]int, length)
	deque := make([]int, 0)
	for i := 0; i < length; i++ {
		for len(deque) != 0 && prices[deque[len(deque)-1]] >= prices[i] {
			ans[deque[len(deque)-1]] = prices[deque[len(deque)-1]] - prices[i]
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		ans[i] = prices[i]
	}
	return ans
}

func main() {
	prices := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			prices = append(prices, n)
		}
	}
	fmt.Println(finalPrices(prices))
}
