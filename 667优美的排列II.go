package main

func constructArray(n int, k int) []int {
	ans := make([]int, n)
	idx := 0
	for i := 0; idx < k; i++ {
		ans[idx] = i + 1
		idx++
		if idx < k {
			ans[idx] = n - i
			idx++
		}
	}
	for i := k; i < n; i++ {
		if k%2 == 0 {
			ans[i] = ans[i-1] - 1
		} else {
			ans[i] = ans[i-1] + 1
		}
	}
	return ans
}
