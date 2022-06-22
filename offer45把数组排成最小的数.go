package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, num := range nums {
		str := strconv.Itoa(num)
		strs = append(strs, str)
	}
	quickSort(strs, 0, len(strs)-1)
	var res strings.Builder
	for _, str := range strs {
		res.WriteString(str)
	}
	return res.String()
}

func quickSort(strs []string, left int, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	tmp := strs[i]
	for i < j {
		con1 := strings.Compare(strs[j]+strs[left], strs[left]+strs[j])
		for (con1 == 1 || con1 == 0) && i < j {
			j--
			con1 = strings.Compare(strs[j]+strs[left], strs[left]+strs[j])
		}
		con2 := strings.Compare(strs[i]+strs[left], strs[left]+strs[i])
		for (con2 == -1 || con2 == 0) && i < j {
			i++
			con2 = strings.Compare(strs[i]+strs[left], strs[left]+strs[i])
		}
		tmp = strs[i]
		strs[i] = strs[j]
		strs[j] = tmp
	}
	strs[i] = strs[left]
	strs[left] = tmp
	quickSort(strs, left, i-1)
	quickSort(strs, i+1, right)
}

func main() {
	arr := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			arr = append(arr, num)
		}
	}
	fmt.Println(minNumber(arr))
}
