package main

import (
	"bufio"
	"fmt"
	"os"
)

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	right := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right < len(s) && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}
		if res < (right - i) {
			res = right - i
		}
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _, _ := reader.ReadLine()
	fmt.Println(lengthOfLongestSubstring(string(s)))
}
