package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//哈希法
//func majorityElement(nums []int) int {
//	count := make(map[int]int)
//	for _, ele := range nums {
//		count[ele]++
//	}
//	for k, v := range count {
//		if v >= (len(nums)+1)/2 {
//			return k
//		}
//	}
//	return 0
//}

//投票法
func majorityElement(nums []int) int {
	votes := 0
	var x int
	for _, ele := range nums {
		if votes == 0 {
			x = ele
		}
		if ele == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}

func main() {
	nums := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strSlice := strings.Split(sca.Text(), ",")
		for _, v := range strSlice {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
	}
	fmt.Println(majorityElement(nums))
}
