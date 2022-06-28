package main

import "fmt"

func translateNum(num int) int {
	a := 1
	b := 1
	y := num % 10
	var x int

	for num != 0 {
		num /= 10
		x = num % 10
		tmp := 10*x + y
		c := a
		if tmp >= 10 && tmp <= 25 {
			c = a + b
		}
		b = a
		a = c
		y = x
	}
	return a
}

//递归法
//func util(nums []int) int {
//	if len(nums) == 1 {
//		return 1
//	}
//	tmp := nums[1]*10 + nums[0]
//	if len(nums) == 2 {
//		if tmp > 25 || tmp < 10 {
//			return 1
//		}
//		return 2
//	}
//	if tmp > 25 || tmp < 10 {
//		return util(nums[1:])
//	}
//	return util(nums[1:]) + util(nums[2:])
//}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(translateNum(num))
}
