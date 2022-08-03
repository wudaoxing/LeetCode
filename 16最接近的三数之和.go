package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func threeSumClosest(nums []int, target int) int {
	ans := nums[0] + nums[1] + nums[2]
	if len(nums) == 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp == target {
				return target
			}
			if abs(tmp-target) < abs(ans-target) {
				ans = tmp
			}
			if tmp < target {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	arr := make([]int, 0)
	var target int
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			arr = append(arr, n)
		}
	}
	fmt.Scan(&target)
	fmt.Println(threeSumClosest(arr, target))
}
