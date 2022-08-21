package main

type MyCircularDeque struct {
	deque []int
	head  int
	tail  int
	nums  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{make([]int, k), -1, 0, 0}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.nums == len(this.deque) {
		return false
	}
	if this.head == -1 {
		this.deque[0] = value
		this.head = 0
		this.tail++
	} else if this.head == 0 {
		this.deque[len(this.deque)-1] = value
		this.head = len(this.deque) - 1
	} else {
		this.head = this.head - 1
		this.deque[this.head] = value
	}
	this.nums++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.nums == len(this.deque) {
		return false
	}
	this.deque[this.tail] = value
	if this.tail == len(this.deque)-1 {
		this.tail = 0
	} else {
		this.tail++
		if this.head == -1 {
			this.head = 0
		}
	}
	this.nums++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.nums == 0 {
		return false
	}
	if this.head == len(this.deque)-1 {
		this.head = 0
	} else {
		this.head++
	}
	this.nums--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.nums == 0 {
		return false
	}
	if this.tail == 0 {
		this.tail = len(this.deque) - 1
	} else {
		this.tail--
	}
	this.nums--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.nums == 0 {
		return -1
	}
	return this.deque[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.nums == 0 {
		return -1
	}
	if this.tail == 0 {
		return this.deque[len(this.deque)-1]
	}
	return this.deque[this.tail-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.nums == 0 {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if this.nums == len(this.deque) {
		return true
	}
	return false
}
