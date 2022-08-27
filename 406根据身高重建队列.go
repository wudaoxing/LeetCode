package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return
}

func main() {
	var n, tmp int
	fmt.Scanf("%d\n", &n)
	people := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		arr := make([]int, 0, 2)
		for j := 0; j < 2; j++ {
			fmt.Scanf("%d", &tmp)
			arr = append(arr, tmp)
		}
		people = append(people, arr)
	}
	fmt.Println(reconstructQueue(people))
}
