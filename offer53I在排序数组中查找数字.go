package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//二分法
func search(nums []int, target int) int {
	return helper(nums, target) - helper(nums, target-1)
}

func helper(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

//func search(nums []int, target int) int {
//	res := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == target {
//			res++
//		}
//		if i-1 >= 0 && nums[i-1] == target && nums[i] != target {
//			break
//		}
//	}
//	return res
//}

func main() {
	fmt.Println("Please input target:")
	var target int
	fmt.Scanf("%d\n", &target)
	fmt.Println("Please input nums:")
	nums := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strArr := strings.Split(sca.Text(), ",")
		for _, v := range strArr {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
	}
	fmt.Println(search(nums, target))
}
