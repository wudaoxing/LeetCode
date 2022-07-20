package main

import "fmt"

func dicesProbability(n int) []float64 {
	res := []float64{1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(res); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += res[j] / 6.0
			}
		}
		res = tmp
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(dicesProbability(n))
}
