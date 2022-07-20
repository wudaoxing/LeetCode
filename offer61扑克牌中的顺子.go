package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func isStraight(nums []int) bool {
//	tmp := make([]int, 14)
//	for i := 0; i < 5; i++ {
//		tmp[nums[i]]++
//	}
//	tag := 0
//	for j := 1; j < 14; j++ {
//		if tmp[j] > 0 {
//			tag = j
//			break
//		}
//	}
//	tagLack := 0
//	//count := 5
//	for k := tag; k <= tag+4; k++ {
//		if k > 13 {
//			tagLack++
//			continue
//		}
//		if tmp[k] > 1 {
//			//count -= tmp[k]
//			//if count == tmp[0] {
//			//	return true
//			//}
//			return false
//		}
//		if tmp[k] == 0 {
//			tagLack++
//		}
//	}
//	if tagLack == tmp[0] {
//		return true
//	}
//	return false
//}

func isStraight(nums []int) bool {
	repeat := make([]int, 14)
	max := 0
	min := 14
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		//_, ok := repeat[num]
		if repeat[num] > 0 {
			return false
		}
		repeat[num]++
	}
	if max-min < 5 {
		return true
	}
	return false
}

func main() {
	nums := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, v := range strs {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
	}
	fmt.Println(isStraight(nums))
}
