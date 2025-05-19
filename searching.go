package main

import "fmt"

const N = 100

func main() {
	var T [N]int
	var x, n, i int

	fmt.Scan(&x, &n)
	i = 0
	for i < n {
		fmt.Scan(&T[i])
		i = i + 1
	}
	fmt.Println(searchBil(T, x, n))

}

func searchBil(T [N]int, x int, n int) int {
	var kanan, kiri, tengah int
	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if T[tengah] == x {
			return tengah
		} else if x < T[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	return -1
}
