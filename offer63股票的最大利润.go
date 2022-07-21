package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	cost := prices[0]
	for i := 1; i < len(prices); i++ {
		if cost > prices[i] {
			cost = prices[i]
		}
		if profit < prices[i]-cost {
			profit = prices[i] - cost
		}
	}
	return profit
}

func main() {
	prices := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, v := range strs {
			n, _ := strconv.Atoi(v)
			prices = append(prices, n)
		}
	}
	fmt.Println(maxProfit(prices))
}
