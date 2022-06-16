package main

import "fmt"

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	dic := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		_, ok := dic[s[i]]
		if ok {
			dic[s[i]] = false
		} else {
			dic[s[i]] = true
		}
	}
	for i := 0; i < len(s); i++ {
		if dic[s[i]] {
			return s[i]
		}
	}
	return ' '
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Printf("%c", firstUniqChar(str))
}
