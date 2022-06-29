package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxValue(grid [][]int) int {
	m := len(grid) - 1
	n := len(grid[0]) - 1
	i := 1
	for ; i <= m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i = 1; i <= n; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}
	j := 1
	for i = 1; i <= m; i++ {
		for j = 1; j <= n; j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}
		}
	}
	return grid[m][n]
}

func main() {
	var m int
	fmt.Scanf("%d\n", &m)
	grid := make([][]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	for i := 0; i < m; i++ {
		if sca.Scan() {
			strSlice := strings.Split(sca.Text(), " ")
			nums := make([]int, 0)
			for _, v := range strSlice {
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}
			grid = append(grid, nums)
		}
	}
	fmt.Println(maxValue(grid))
}
