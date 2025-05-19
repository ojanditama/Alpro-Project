package main

const NMAX int = 100

type tabInt [NMAX]int

func main() {
	var A tabInt
	var N int

}

func urut(A *tabInt, N int) {
	var pass, idx, i, temp int

	pass = 1

	for pass <= N-1 {
		idx = pass - 1
		i = pass

		for i < N {
			if A[idx] < A[i] {
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
