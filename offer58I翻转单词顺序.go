package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func reverseWords(s string) string {
	var bt bytes.Buffer
	strs := strings.Fields(s)
	for i := len(strs) - 1; i >= 0; i-- {
		bt.WriteString(strs[i])
		if i != 0 {
			bt.WriteString(" ")
		}
	}
	return bt.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := reader.ReadLine()
	fmt.Println(reverseWords(string(res)))
}
