package main

func uniqueLetterString(s string) (ans int) {
	chars := map[rune][]int{}
	for i, c := range s {
		chars[c] = append(chars[c], i)
	}
	for _, arr := range chars {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return ans
}
