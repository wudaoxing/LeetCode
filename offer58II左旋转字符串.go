package main

import (
	"fmt"
)

//func reverseLeftWords(s string, n int) string {
//	var res strings.Builder
//	for i := n; i < len(s); i++ {
//		res.WriteString(string(s[i]))
//	}
//	for j := 0; j < n; j++ {
//		res.WriteString(string(s[j]))
//	}
//	return res.String()
//}

//func reverseLeftWords(s string, n int) string {
//	res := make([]byte, 0, len(s)-1)
//	for i := n; i < n+len(s); i++ {
//		res = append(res, s[i%len(s)])
//	}
//	return string(res)
//}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func main() {
	var s string
	var n int
	fmt.Scanf("%s %d", &s, &n)
	fmt.Println(reverseLeftWords(s, n))
}
