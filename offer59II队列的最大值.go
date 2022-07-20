package main

import "fmt"

type MaxQueue struct {
	arr []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{arr: make([]int, 0),
		max: make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.max) > 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
	this.arr = append(this.arr, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.arr) == 0 {
		return -1
	}
	if this.arr[0] == this.max[0] {
		this.max = this.max[1:]
	}
	tmp := this.arr[0]
	this.arr = this.arr[1:]
	return tmp
}

func main() {
	obj := Constructor()
	obj.Push_back(94)
	obj.Push_back(16)
	obj.Push_back(89)
	obj.Pop_front()
	obj.Push_back(22)
	fmt.Println(obj.Pop_front())
}
