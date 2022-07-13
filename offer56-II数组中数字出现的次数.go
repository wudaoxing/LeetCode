package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func singleNumber(nums []int) int {
	two := 0
	one := 0
	for _, num := range nums {
		one = one ^ num&(^two)
		two = two ^ num&(^one)
	}
	return one
}

func main() {
	nums := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strArr := strings.Split(sca.Text(), ",")
		for _, v := range strArr {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
	}
	fmt.Println(singleNumber(nums))
}
