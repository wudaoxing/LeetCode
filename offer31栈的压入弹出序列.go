package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	prePush := 0
	prePop := 0
	for prePush < len(pushed) {
		stack = append(stack, pushed[prePush])
		for len(stack) > 0 && stack[len(stack)-1] == popped[prePop] {
			stack = stack[:len(stack)-1]
			prePop++
		}
		prePush++
	}
	return len(stack) == 0
}

func main() {
	pushed := make([]int, 0, 5)
	popped := make([]int, 0, 5)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strSlice := strings.Split(sca.Text(), ",")
		for _, v := range strSlice {
			num, _ := strconv.Atoi(v)
			pushed = append(pushed, num)
		}
	}
	if sca.Scan() {
		strSlice := strings.Split(sca.Text(), ",")
		for _, v := range strSlice {
			num, _ := strconv.Atoi(v)
			popped = append(popped, num)
		}
	}
	fmt.Println(validateStackSequences(pushed, popped))

}
