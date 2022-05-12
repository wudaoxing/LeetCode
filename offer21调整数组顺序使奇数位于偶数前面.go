package main

import "fmt"

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < len(nums) && nums[left]%2 != 0 {
			left++
		}
		for right >= 0 && nums[right]%2 == 0 {
			right--
		}
		if left < right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			right--
			left++
		}
	}
	return nums
}

func main() {
	nums := []int{1, 3, 5}
	fmt.Println(exchange(nums))
}
