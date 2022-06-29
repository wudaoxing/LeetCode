package main

import "fmt"

//动态规划+哈希表
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	tmp := 0
	dic := map[byte]int{}
	for j := 0; j < len(s); j++ {
		i, ok := dic[s[j]]
		if !ok {
			i = -1
		}
		dic[s[j]] = j
		if tmp < j-i {
			tmp++
		} else {
			tmp = j - i
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

//动态规划+线性遍历
//func lengthOfLongestSubstring(s string) int {
//	res := 0
//	if (len(s) == 0) || (s == "") {
//		return 0
//	}
//	tmp := 0
//	for i := 0; i < len(s); i++ {
//		j := i - 1
//		for ; j >= 0; j-- {
//			if s[i] == s[j] {
//				break
//			}
//		}
//		if tmp < i-j {
//			tmp++
//		} else {
//			tmp = i - j
//		}
//		if tmp > res {
//			res = tmp
//		}
//	}
//	return res
//}

func main() {
	var s string
	fmt.Scanf("%[^\n]", &s)
	fmt.Println(lengthOfLongestSubstring(s))
}
