package main

import "fmt"

type MedianFinder struct {
	A []int
}

func Constructor() MedianFinder {
	return MedianFinder{make([]int, 0)}
}

/*堆*/
//type MedianFinder struct {
//	A []int //小顶堆，存储较大的数
//	B []int //大顶堆，存储较小的数
//}
//
//func Constructor() MedianFinder {
//	return MedianFinder{
//		make([]int, 0),
//		make([]int, 0),
//	}
//}
//func buildMinHeap(arr []int) {
//	for i := (len(arr) - 2) / 2; i >= 0; i-- {
//		minHeapify(arr, i)
//	}
//}
//func minHeapify(arr []int, i int) {
//	left := i*2 + 1
//	right := i*2 + 2
//	min := i
//	if left < len(arr) && arr[left] < arr[min] {
//		min = left
//	}
//	if right < len(arr) && arr[right] < arr[min] {
//		min = right
//	}
//	if min != i {
//		temp := arr[i]
//		arr[i] = arr[min]
//		arr[min] = temp
//		minHeapify(arr, min)
//	}
//}
//
//func buildMaxHeap(arr []int) {
//	for i := (len(arr) - 2) / 2; i >= 0; i-- {
//		maxHeapify(arr, i)
//	}
//}
//func maxHeapify(arr []int, i int) {
//	left := i*2 + 1
//	right := i*2 + 2
//	max := i
//	if left < len(arr) && arr[left] > arr[max] {
//		max = left
//	}
//	if right < len(arr) && arr[right] > arr[max] {
//		max = right
//	}
//	if max != i {
//		temp := arr[i]
//		arr[i] = arr[max]
//		arr[max] = temp
//		minHeapify(arr, max)
//	}
//}
//
//func (this *MedianFinder) AddNum(num int) {
//	if len(this.A) == len(this.B) {
//		//向B中插入元素
//		this.B = append(this.B, num)
//		buildMaxHeap(this.B)
//		//将B的堆顶元素放入A中
//		this.A = append(this.A, this.B[0])
//		this.B = this.B[1:]
//		buildMinHeap(this.A)
//		buildMaxHeap(this.B)
//	} else {
//		//向A中插入元素
//		this.A = append(this.A, num)
//		buildMinHeap(this.A)
//		//将A的堆顶元素放入B中
//		this.B = append(this.B, this.A[0])
//		this.A = this.A[1:]
//		buildMinHeap(this.A)
//		buildMaxHeap(this.B)
//	}
//}
//
//func (this *MedianFinder) FindMedian() float64 {
//	if len(this.A) == len(this.B) {
//		return float64(this.A[0]+this.B[0]) / 2.0
//	} else {
//		return float64(this.A[0])
//	}
//}

func (this *MedianFinder) AddNum(num int) {
	if len(this.A) == 0 {
		this.A = append(this.A, num)
	} else {
		//二分查找
		left := 0
		right := len(this.A) - 1
		for left <= right {
			mid := left + (right-left)/2
			if this.A[mid] <= num {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		//插入元素
		this.A = append(this.A, num)
		for i := len(this.A) - 1; i > left; i-- {
			temp := this.A[i]
			this.A[i] = this.A[i-1]
			this.A[i-1] = temp
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	len := len(this.A)
	if len%2 == 0 {
		return float64(this.A[len/2]+this.A[len/2-1]) / 2
	} else {
		return float64(this.A[len/2])
	}
}
func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
}
