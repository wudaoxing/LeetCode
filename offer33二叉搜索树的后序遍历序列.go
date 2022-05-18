package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func verifyPostorder(postorder []int) bool {
	stack := make([]int, 0)
	root := math.MaxInt64
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) != 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}

func main() {
	postorder := make([]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		temp := strings.Split(sca.Text(), ",")
		for _, v := range temp {
			val, _ := strconv.Atoi(v)
			postorder = append(postorder, val)
		}
	}
	fmt.Println(verifyPostorder(postorder))
}
