package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func singleNumbers(nums []int) []int {
	x := 0
	y := 0
	n := 0
	m := 1
	for _, num := range nums {
		n ^= num
	}
	for n&m == 0 {
		m <<= 1
	}
	for _, num := range nums {
		if num&m != 0 {
			x ^= num
		} else {
			y ^= num
		}
	}
	return []int{x, y}
}

func main() {
	arr := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strArr := strings.Split(sca.Text(), ",")
		for _, v := range strArr {
			num, _ := strconv.Atoi(v)
			arr = append(arr, num)
		}
	}
	fmt.Println(singleNumbers(arr))
}
