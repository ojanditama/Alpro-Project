package main

import "fmt"

const NMAX int = 99

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int

	fmt.Scan(&nData)
	bacaData(&data, nData)
	cetakData(data, nData)
	fmt.Print("\n")
	selectionSort(&data, nData)
	cetakData(data, nData)
}

func selectionSort(A *tabInt, N int) {
	var i, idx, pass int
	var temp int

	pass = 1

	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			if A[i] > A[idx] {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func bacaData(A *tabInt, N int) {
	var i int

	for i = 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, N int) {
	var i int

	for i = 0; i < N; i++ {
		fmt.Printf("%d ", A[i])
	}
}
