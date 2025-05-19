// package main

// import "fmt"

// const N int = 1000

// func main() {
// 	var T [N]int
// 	var jumlah int
// 	fmt.Scanln(&jumlah)
// 	for i := 0; i < jumlah; i++ {
// 		fmt.Scan(&T[i])
// 	}
// 	fmt.Println(ganjilTerbesar(T, jumlah))
// }

// func ganjilTerbesar(T [N]int, jumlah int) int {
// 	var i, maxGanjil int

// 	for i = 0; i < jumlah; i++ {
// 		if T[i]%2 != 0 && T[i] > maxGanjil {
// 			maxGanjil = T[i]
// 		}
// 	}
// 	return maxGanjil
// }
//============================GANJIL TERBESAR===========================//

package main

import "fmt"

type arrOfInt [N]int

const N int = 1000

func main() {
	var T, P arrOfInt
	var total, n int
	var length int = 0

	fmt.Scan(&total)
	for j := 0; j < total; j++ {
		fmt.Scan(&T[j])
	}
	prima(T, total, &length, &P)
	for n < length {
		fmt.Print(P[n], " ")
		n++
	}
}

func prima(T arrOfInt, total int, length *int, P *arrOfInt) {
	var m int
	for i := 0; i < total; i++ {
		m = 0
		for k := 1; k <= T[i]; k++ {
			if T[i]%k == 0 {
				m++
			}
		}
		if m == 2 { // Bilangan prima hanya punya 2 faktor: 1 dan dirinya sendiri
			P[*length] = T[i]
			*length++
		}
	}
}
