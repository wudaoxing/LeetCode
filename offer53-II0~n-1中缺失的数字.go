package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
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
	fmt.Println(missingNumber(nums))
}
