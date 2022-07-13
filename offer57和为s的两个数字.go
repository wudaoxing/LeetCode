package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		s := nums[i] + nums[j]
		if s == target {
			return []int{nums[i], nums[j]}
		} else if s < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

func main() {
	var target int
	fmt.Print("target=")
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
	fmt.Println(twoSum(nums, target))
}
