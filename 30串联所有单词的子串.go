package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	wordNum := len(words)
	if wordNum == 0 {
		return ans
	}
	wordLen := len(words[0])
	//allWords存储words中的单词及其频数
	allWords := map[string]int{}
	for _, w := range words {
		allWords[w]++
	}
	//将滑动窗口的移动分为wordLen类情况
	for j := 0; j < wordLen; j++ {
		hasWords := map[string]int{}
		num := 0
		//每次移动一个单词
		for i := j; i < len(s)-wordNum*wordLen+1; i = i + wordLen {
			hasRemoved := false
			//判断当前子串是否匹配
			for num < wordNum {
				word := s[i+num*wordLen : i+(num+1)*wordLen]
				_, ok := allWords[word]
				if ok {
					hasWords[word]++
					if hasWords[word] > allWords[word] {
						//判断（符合但次数超过）超过的单词是否移除的标志
						hasRemoved = true
						//移除单词个数
						removeNum := 0
						for hasWords[word] > allWords[word] {
							firstWord := s[i+removeNum*wordLen : i+(removeNum+1)*wordLen]
							hasWords[firstWord]--
							removeNum++
						}
						num = num - removeNum + 1
						i = i + (removeNum-1)*wordLen
						break
					}
				} else {
					hasWords = map[string]int{}
					i = i + num*wordLen
					num = 0
					break
				}
				num++
			}
			if num == wordNum {
				ans = append(ans, i)
			}
			//子串完全匹配
			if num > 0 && !hasRemoved {
				firstWord := s[i : i+wordLen]
				hasWords[firstWord]--
				num = num - 1
			}
		}
	}
	return ans
}

func main() {
	var s string
	words := make([]string, 0)
	fmt.Print("s=")
	fmt.Scanf("%s\n", &s)
	fmt.Print("words=")
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		words = strings.Split(sca.Text(), ",")
	}
	fmt.Println(findSubstring(s, words))
}
