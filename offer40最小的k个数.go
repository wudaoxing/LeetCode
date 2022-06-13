package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*堆排序*/
//func heapSort(arr []int) {
//	if arr == nil || len(arr) == 0 {
//		return
//	}
//	len := len(arr)
//	//构建大顶堆
//	buildMaxHeap(arr, len)
//}
//
//func buildMaxHeap(arr []int, len int) {
//	for i := (len - 1) / 2; i >= 0; i-- {
//		heapify(arr, i, len)
//	}
//}
//
//func heapify(arr []int, i int, len int) {
//	left := 2*i + 1
//	right := 2*i + 2
//	max := i
//	if left < len && arr[left] > arr[max] {
//		max = left
//	}
//	if right < len && arr[right] > arr[max] {
//		max = right
//	}
//	if max != i {
//		swap(arr, i, max)
//		heapify(arr, max, len)
//	}
//}
//
//func swap(arr []int, i int, j int) {
//	temp := arr[i]
//	arr[i] = arr[j]
//	arr[j] = temp
//}
//
//func getLeastNumbers(arr []int, k int) []int {
//	if arr == nil || len(arr) == 0 || k == 0{
//		return nil
//	}
//	res := arr[:k]
//	heapSort(res)
//	for _, ele := range arr[k:] {
//		if ele < res[0] {
//			res = res[1:]
//			res = append(res, ele)
//			heapSort(res)
//		}
//	}
//	return res
//}

/*快速排序*/
func quickSort(arr []int, low int, high int, k int) []int {
	j := partition(arr, low, high)
	if j == k {
		return arr[:j+1]
	} else if j > k {
		return quickSort(arr, low, j-1, k)
	} else {
		return quickSort(arr, j+1, high, k)
	}
}

func partition(arr []int, low int, high int) int {
	v := arr[low]
	i, j := low+1, high
	for i < j {
		for i <= j && arr[j] >= v {
			j--
		}
		for i <= j && arr[i] <= v {
			i++
		}
		if i >= j {
			break
		}
		t := arr[i]
		arr[i] = arr[j]
		arr[j] = t
	}
	arr[low] = arr[j]
	arr[j] = v
	return j
}

func getLeastNumbers(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 || k == 0 {
		return nil
	}

	return quickSort(arr, 0, len(arr)-1, k-1)
}

func main() {
	arr := make([]int, 0)
	k := 0
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		stringK := sca.Text()
		k, _ = strconv.Atoi(stringK)
	}
	if sca.Scan() {
		stringArr := strings.Split(sca.Text(), ",")
		for _, ele := range stringArr {
			num, _ := strconv.Atoi(ele)
			arr = append(arr, num)
		}
	}
	fmt.Println(getLeastNumbers(arr, k))
}
