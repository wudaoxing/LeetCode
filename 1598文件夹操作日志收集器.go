package main

func minOperations(logs []string) int {
	init := 0
	for _, s := range logs {
		if s == "../" {
			if init != 0 {
				init--
			}
		} else if s == "./" {
			continue
		} else {
			init++
		}
	}
	return init
}
