package main

import "fmt"

const N int = 1000

func main() {
	var T [N]int
	var total, j int

	fmt.Scan(&total)
	for j = 0; j < total; j++ {
		fmt.Scan(&T[j])
	}
	fmt.Print(yangTertukar(T, total))
}

func yangTertukar(T [N]int, total int) int {
	var jumlahTertukar, i int

	for i = 0; i < total-1; i++ {
		if T[i] > T[i+1] {
			jumlahTertukar++
		}
	}
	return jumlahTertukar
}
