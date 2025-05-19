package main

import "fmt"

type pemain struct {
	poin int
}

const NMAX int = 1024

type dataPemain [NMAX]pemain

func main() {
	var himpunan dataPemain
	var n int

	fmt.Scanln(&n)
	isiArray(&himpunan, n)
	selectionSort(&himpunan, n)
	showArray(himpunan, n)
}

func isiArray(himpunan *dataPemain, n int) {
	for i := 0; i < n; i++ {
		fmt.Scanln(&himpunan[i].poin)
	}
}

func selectionSort(himpunan *dataPemain, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if himpunan[minIdx].poin > himpunan[j].poin {
				minIdx = j
			}
		}
		// tukar
		himpunan[i], himpunan[minIdx] = himpunan[minIdx], himpunan[i]
	}
}

func showArray(himpunan dataPemain, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].poin)
	}
}
