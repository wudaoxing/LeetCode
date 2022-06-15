package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}

func main() {
	nums := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		str_nums := strings.Split(sca.Text(), ",")
		for _, ele := range str_nums {
			num, _ := strconv.Atoi(ele)
			nums = append(nums, num)
		}
	}
	fmt.Println(maxSubArray(nums))
}
