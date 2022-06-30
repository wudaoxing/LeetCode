package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reversePairs(nums []int) int {
	tmp := make([]int, len(nums))
	return mergeSort(nums, tmp, 0, len(nums)-1)
}

func mergeSort(nums []int, tmp []int, left int, right int) int {
	if left >= right {
		return 0
	}
	m := left + (right-left)/2
	res := mergeSort(nums, tmp, left, m) + mergeSort(nums, tmp, m+1, right)
	i := left
	j := m + 1
	for k := left; k <= right; k++ {
		tmp[k] = nums[k]
	}
	for k := left; k <= right; k++ {
		if i == m+1 {
			nums[k] = tmp[j]
			j++
		} else if j == right+1 || tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
			res += m - i + 1
		}

	}
	return res
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
	fmt.Println(reversePairs(nums))
}
