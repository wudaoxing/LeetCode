package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func constructArr(a []int) []int {
	if len(a) == 0 {
		return a
	}
	b := make([]int, len(a))
	b[0] = 1
	tmp := 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	for j := len(a) - 2; j >= 0; j-- {
		tmp *= a[j+1]
		b[j] *= tmp
	}
	return b
}

func main() {
	a := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, v := range strs {
			n, _ := strconv.Atoi(v)
			a = append(a, n)
		}
	}
	fmt.Println(constructArr(a))
}
