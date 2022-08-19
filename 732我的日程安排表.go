package main

import (
	"fmt"
)

/*差分数组*/
/*
type MyCalendarThree struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{redblacktree.NewWithIntComparator()}
}

func (t MyCalendarThree) add(x, delta int) {
	if val, ok := t.Get(x); ok {
		delta += val.(int)
	}
	t.Put(x, delta)
}

func (t MyCalendarThree) Book(start, end int) (ans int) {
	t.add(start, 1)
	t.add(end, -1)

	maxBook := 0
	for it := t.Iterator(); it.Next(); {
		maxBook += it.Value().(int)
		if maxBook > ans {
			ans = maxBook
		}
	}
	return
}*/

/*线段树*/
type pair struct {
	num  int
	lazy int
}

type MyCalendarThree map[int]pair

func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

func (t MyCalendarThree) update(start, end, l, r, idx int) {
	if r < start || end < l {
		return
	}
	if start <= l && r <= end {
		p := t[idx]
		p.num++
		p.lazy++
		t[idx] = p
	} else {
		mid := (l + r) / 2
		t.update(start, end, l, mid, idx*2)
		t.update(start, end, mid+1, r, idx*2+1)
		p := t[idx]
		p.num = p.lazy + max(t[idx*2].num, t[idx*2+1].num)
		t[idx] = p
	}
}

func (t MyCalendarThree) Book(start, end int) int {
	t.update(start, end-1, 0, 1e9, 1)
	return t[1].num
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	obj := Constructor()
	param_1 := obj.Book(10, 20)
	param_2 := obj.Book(50, 60)
	param_3 := obj.Book(10, 40)
	param_4 := obj.Book(5, 15)
	param_5 := obj.Book(5, 10)
	param_6 := obj.Book(25, 55)
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)
}
