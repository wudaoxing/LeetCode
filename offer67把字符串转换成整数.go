package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	start := -1
	sign := 1
	res := 0
	bndry := math.MaxInt32 / 10
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			start = i
			break
		}
		if i == len(str)-1 {
			return 0
		}
	}
	if str[start] == '-' {
		sign = -1
	}
	if str[start] == '+' || str[start] == '-' {
		start++
	}
	for j := start; j < len(str); j++ {
		if str[j] < '0' || str[j] > '9' {
			break
		}
		if (res > bndry) || (res == bndry && str[j] > '7') {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		res = res*10 + int(str[j]-'0')
	}
	return sign * res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()
	fmt.Println(strToInt(string(str)))
}
