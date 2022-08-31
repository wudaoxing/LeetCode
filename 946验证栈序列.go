package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateStackSequences(pushed []int, popped []int) bool {
	idx := 0
	j := 0
	for i := 0; i < len(pushed); i++ {
		pushed[idx] = pushed[i]
		idx++
		for idx > 0 && pushed[idx-1] == popped[j] {
			idx--
			j++
		}
	}
	return idx == 0
}

func main() {
	pushed := []int{}
	popped := []int{}
	sca := bufio.NewScanner(os.Stdin)
	fmt.Print("pushed=")
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			pushed = append(pushed, n)
		}
	}
	fmt.Print("popped=")
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			popped = append(popped, n)
		}
	}
	fmt.Println(validateStackSequences(pushed, popped))
}
