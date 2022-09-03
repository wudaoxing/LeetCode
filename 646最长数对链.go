package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	cur := math.MinInt
	ans := 0
	for _, p := range pairs {
		if cur < p[0] {
			ans++
			cur = p[1]
		}
	}
	return ans
}

func main() {
	pairs := make([][]int, 0)
	sca := bufio.NewScanner(os.Stdin)
	if sca.Scan() {
		strs := strings.Split(sca.Text(), ",")
		for i := 0; i < len(strs); {
			n1, _ := strconv.Atoi(strs[i])
			n2, _ := strconv.Atoi(strs[i+1])
			pairs = append(pairs, []int{n1, n2})
			i += 2
		}
	}
	fmt.Println(findLongestChain(pairs))
}
