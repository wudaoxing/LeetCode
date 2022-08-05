package main

import "fmt"

type MyCalendarThree struct {
	Calendar map[[2]int]int
	MaxBook  int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{make(map[[2]int]int, 0), 0}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	for k, _ := range this.Calendar {
		if start >= k[1] || end <= k[0] {
			continue
		} else {
			this.Calendar[k]++
		}
	}
	this.Calendar[[2]int{start, end}]++
	for _, v := range this.Calendar {
		if v > this.MaxBook {
			this.MaxBook = v
		}
	}
	return this.MaxBook
}

func main() {
	obj := Constructor()
	param_1 := obj.Book(24, 40)
	param_2 := obj.Book(43, 50)
	param_3 := obj.Book(27, 43)
	param_4 := obj.Book(5, 21)
	param_5 := obj.Book(30, 40)
	param_6 := obj.Book(14, 29)
	param_7 := obj.Book(3, 19)
	param_8 := obj.Book(3, 14)
	param_9 := obj.Book(25, 39)
	param_10 := obj.Book(6, 19)
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)
	fmt.Println(param_7)
	fmt.Println(param_8)
	fmt.Println(param_9)
	fmt.Println(param_10)
}
