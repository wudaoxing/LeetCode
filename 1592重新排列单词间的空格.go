package main

func reorderSpaces(text string) string {
	countSpaces := 0
	strs := make([]string, 0)
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			countSpaces++
		}
		if text[i] != ' ' && (i == 0 || text[i-1] == ' ') {
			tmp := make([]byte, 0)
			for ; i < len(text) && text[i] != ' '; i++ {
				tmp = append(tmp, text[i])
			}
			strs = append(strs, string(tmp))
			i--
		}
	}
	var num int
	var remainder int
	if len(strs) == 1 {
		num = 0
		remainder = countSpaces
	} else {
		num = countSpaces / (len(strs) - 1)
		remainder = countSpaces % (len(strs) - 1)
	}
	tmp1 := make([]byte, 0)
	for num > 0 {
		tmp1 = append(tmp1, ' ')
		num--
	}
	tmp2 := make([]byte, 0)
	for remainder > 0 {
		tmp2 = append(tmp2, ' ')
		remainder--
	}
	for i := 1; i < len(strs); i++ {
		strs[0] = strs[0] + string(tmp1) + strs[i]
	}
	strs[0] = strs[0] + string(tmp2)
	return strs[0]
}
