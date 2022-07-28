package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, s := range strs {
		cnt := [26]int{}
		for _, c := range s {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], s)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := make([]string, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs = strings.Split(sca.Text(), ",")
	}
	fmt.Println(groupAnagrams(strs))
}
