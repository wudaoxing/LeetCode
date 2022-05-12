package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || matrix == nil {
		return nil
	}
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	x := 0
	res := make([]int, (right+1)*(bottom+1))
	for {
		for i := left; i <= right; i++ {
			res[x] = matrix[top][i]
			x++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[x] = matrix[i][right]
			x++
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			res[x] = matrix[bottom][i]
			x++
		}
		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			res[x] = matrix[i][left]
			x++
		}
		left++
		if left > right {
			break
		}
	}

	return res
}

func main() {
	var arr, col int //输入矩阵的行列数
	fmt.Println("请输入矩阵行数：")
	fmt.Scan(&arr)
	fmt.Println("请输入矩阵列数：")
	fmt.Scan(&col)
	matrix := make([][]int, arr)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, col)
	}
	fmt.Println("请输入矩阵：")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	res := spiralOrder(matrix)
	fmt.Println(res)

}
