package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//模拟
//func numSpecial(mat [][]int) (ans int) {
//	rows := len(mat)
//	cols := len(mat[0])
//	r := make([]int, rows)
//	c := make([]int, cols)
//	for i := 0; i < rows; i++ {
//		for j := 0; j < cols; j++ {
//			r[i] += mat[i][j]
//			c[j] += mat[i][j]
//		}
//	}
//	for i := 0; i < rows; i++ {
//		for j := 0; j < cols; j++ {
//			if mat[i][j] == 1 && r[i] == 1 && c[j] == 1 {
//				ans++
//			}
//		}
//	}
//	return
//}

//列的标记值
func numSpecial(mat [][]int) (ans int) {
	for i, row := range mat {
		tmp := 0
		for _, n := range row {
			tmp += n
		}
		if i == 0 {
			tmp--
		}
		if tmp > 0 {
			for j, n := range row {
				if n == 1 {
					mat[0][j] += tmp
				}
			}
		}
	}
	for _, n := range mat[0] {
		if n == 1 {
			ans++
		}
	}
	return
}

func main() {
	var rows int
	mat := make([][]int, 0)
	fmt.Scanf("%d\n", &rows)
	sca := bufio.NewScanner(os.Stdin)
	for i := 0; i < rows; i++ {
		if sca.Scan() {
			nums := make([]int, 0)
			strs := strings.Split(sca.Text(), ",")
			for _, s := range strs {
				n, _ := strconv.Atoi(s)
				nums = append(nums, n)
			}
			mat = append(mat, nums)
		}
	}
	fmt.Println(numSpecial(mat))
}
