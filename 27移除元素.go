package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeElement(nums []int, val int) int {
	len := len(nums)
	if (len == 0) || (len == 1 && nums[0] == val) {
		return 0
	}
	left := 0
	right := len - 1
	for left <= right {
		for left < len && nums[left] != val {
			left++
		}
		for right >= 0 && nums[right] == val {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			nums[right] = val
			left++
			right--
		}
	}
	return right + 1
}

func main() {
	nums := make([]int, 0)
	var val int
	fmt.Print("nums=")
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
	}
	fmt.Print("val=")
	fmt.Scan(&val)
	fmt.Println(removeElement(nums, val))
}
