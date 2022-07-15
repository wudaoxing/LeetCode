package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func maxSlidingWindow(nums []int, k int) []int {
//	if len(nums) == 0 {
//		return nums
//	}
//	max := math.MinInt
//	res := make([]int, 0, len(nums)-k+1)
//	for i := 0; i < k; i++ {
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	res = append(res, max)
//	for j := k; j < len(nums); j++ {
//		if nums[j-k] != max {
//			if nums[j] > max {
//				max = nums[j]
//			}
//		} else {
//			max = nums[j-k+1]
//			for g := j - k + 2; g <= j; g++ {
//				if nums[g] > max {
//					max = nums[g]
//				}
//			}
//		}
//		res = append(res, max)
//	}
//	return res
//}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return make([]int, 0)
	}
	res := make([]int, 0, len(nums)-k+1)
	queue := make([]int, 0, k)
	for i := 0; i < k; i++ {
		for len(queue) != 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res = append(res, queue[0])
	for j := k; j < len(nums); j++ {
		if queue[0] == nums[j-k] {
			queue = queue[1:]
		}
		for len(queue) != 0 && queue[len(queue)-1] < nums[j] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[j])
		res = append(res, queue[0])
	}
	return res
}

func main() {
	var k int
	fmt.Scanf("%d\n", &k)
	nums := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, v := range strs {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
	}
	fmt.Println(maxSlidingWindow(nums, k))
}
