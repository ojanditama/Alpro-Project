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
	InsertionSort(&data, nData)
	cetakData(data, nData)
}

func InsertionSort(A *tabInt, N int) {
	var i, pass, temp int

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp < A[i-1] {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
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
	fmt.Println()
}
