package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX = 100

type komentar struct {
	usn      string
	teks     string
	sentimen string
}

var daftarKomentar [NMAX]komentar
var jumlahKomentar int

var komentarPositif = [NMAX]string{"baik", "bagus", "mantap", "hebat"}
var komentarNegatif = [NMAX]string{"jelek", "buruk", "parah", "bodoh"}

func analisisSentimen(teks string) string {
	var i int

	for i = 0; i < len(komentarNegatif); i++ {
		if cari(komentarNegatif[i], teks) {
			return "negatif"
		}
	}

	for i = 0; i < len(komentarPositif); i++ {
		if cari(komentarPositif[i], teks) {
			return "positif"
		}
	}

	return "netral"
}

func masukKomentar() {
	var usn, teks string
	for {

		fmt.Print("\nMasukkan ID (atau # untuk keluar): ")
		fmt.Scan(&usn)
		if usn == "#" {
			break
		}
		fmt.Print("Masukkan Komentar (atau # untuk keluar): ")
		fmt.Scan(&teks)
		if teks == "#" {
			break
		}
		if jumlahKomentar < NMAX {
			var dataKomentar komentar
			dataKomentar.usn = usn
			dataKomentar.teks = teks
			dataKomentar.sentimen = analisisSentimen(teks)

			daftarKomentar[jumlahKomentar] = dataKomentar
			jumlahKomentar++
		} else {
			fmt.Println("Data komentar penuh.")
			break
		}
	}
}

func tampilKomentar() {
	var i int
	fmt.Println("\nDaftar Komentar:")
	for i = 0; i < jumlahKomentar; i++ {
		fmt.Printf("[%s] %s => Sentimen: %s\n", daftarKomentar[i].usn, daftarKomentar[i].teks, daftarKomentar[i].sentimen)
	}
}

func statistikSentimen() {
	var i int
	var positif, netral, negatif int
	for i = 0; i < jumlahKomentar; i++ {
		if daftarKomentar[i].sentimen == "positif" {
			positif++
		} else if daftarKomentar[i].sentimen == "netral" {
			netral++
		} else if daftarKomentar[i].sentimen == "negatif" {
			negatif++
		}
	}
	fmt.Println("\nStatistik Sentimen:")
	fmt.Println("Positif:", positif)
	fmt.Println("Netral :", netral)
	fmt.Println("Negatif:", negatif)
}

func ubahKomentar() {
	var usn, komentarBaru string
	var i int
	fmt.Print("Masukkan ID komentar yang ingin diubah: ")
	fmt.Scan(&usn)

	for i = 0; i < jumlahKomentar; i++ {
		if daftarKomentar[i].usn == usn {
			fmt.Printf("Komentar lama: %s\n", daftarKomentar[i].teks)
			fmt.Print("Masukkan komentar baru: ")
			fmt.Scan(&komentarBaru)

			daftarKomentar[i].teks = komentarBaru
			daftarKomentar[i].sentimen = analisisSentimen(komentarBaru)

			fmt.Println("Komentar berhasil diubah.")
			return
		}
	}
	fmt.Println("Komentar tidak ditemukan.")
}

func hapusKomentar() {
	var i, j int
	var id string
	fmt.Print("Masukkan ID komentar yang ingin dihapus: ")
	fmt.Scan(&id)
	for i = 0; i < jumlahKomentar; i++ {
		if daftarKomentar[i].usn == id {
			for j = i; j < jumlahKomentar-1; j++ {
				daftarKomentar[j] = daftarKomentar[j+1]
			}
			jumlahKomentar--
			fmt.Println("Komentar berhasil dihapus.")
			return
		}
	}
	fmt.Println("Komentar tidak ditemukan.")
}

func sequentialSearch(keyword string) {
	var found bool
	var i int

	fmt.Println("\nHasil Pencarian (Sequential Search):")
	found = false
	for i = 0; i < jumlahKomentar; i++ {
		if cari(daftarKomentar[i].teks, keyword) {
			fmt.Printf("[%s] %s => %s\n", daftarKomentar[i].usn, daftarKomentar[i].teks, daftarKomentar[i].sentimen)
			found = true
		}
	}
	if !found {
		fmt.Println("Komentar tidak ditemukan.")
	}
}

func selectionSortKomentar() {
	var i, j, minIdx int
	for i = 0; i < jumlahKomentar-1; i++ {
		minIdx = i
		for j = i + 1; j < jumlahKomentar; j++ {
			if daftarKomentar[j].teks < daftarKomentar[minIdx].teks {
				minIdx = j
			}
		}
		if minIdx != i {
			daftarKomentar[i], daftarKomentar[minIdx] = daftarKomentar[minIdx], daftarKomentar[i]
		}
	}
}

func binarySearch(keyword string) {
	var low, mid, high int
	var found bool
	var teks string

	selectionSortKomentar()

	low = 0
	high = jumlahKomentar - 1

	found = false
	for low <= high {
		mid = (low + high) / 2
		teks = daftarKomentar[mid].teks

		if cari(teks, keyword) {
			fmt.Printf("[%s] %s => %s\n", daftarKomentar[mid].usn, daftarKomentar[mid].teks, daftarKomentar[mid].sentimen)
			found = true
			break
		} else if teks < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Println("Komentar tidak ditemukan.")
	}
}

func color(text string) {
	cmd := exec.Command("cmd", "/c", "color", text)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func box() {
	clear()
	color("a")
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                   Aplikasi Analisis Sentimen                 ║")
	fmt.Println("║                  Created by : Fauzan & Mario                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                              ║")
	fmt.Println("║                           MAIN MENU                          ║")
	fmt.Println("║                                                              ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║     Selamat datang di aplikasi analisis komentar online.     ║")
	fmt.Println("║                                                              ║")
	fmt.Println("║     1. Masukkan Komentar                                     ║")
	fmt.Println("║     2. Tampilkan Komentar                                    ║")
	fmt.Println("║     3. Statistik Sentimen                                    ║")
	fmt.Println("║     4. Ubah Komentar                                         ║")
	fmt.Println("║     5. Hapus Komentar                                        ║")
	fmt.Println("║     6. Cari Komentar (Sequential Search)                     ║")
	fmt.Println("║     7. Cari Komentar (Binary Search)                         ║")
	fmt.Println("║     8. Urutkan Komentar (Selection Sort berdasarkan teks)    ║")
	fmt.Println("║     9. Keluar                                                ║")
	fmt.Println("║                                                              ║")
	fmt.Println("║     Silakan pilih menu yang Anda butuhkan.                   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Print("SELECT> ")
}

func subBox(judul string) {
	clear()
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Printf("║ %-60s ║\n", judul)
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
}

func main() {
	var pilih int
	for pilih != 9 {
		box()
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			subBox("Masukkan Komentar")
			masukKomentar()
		case 2:
			subBox("Tampilkan Komentar")
			tampilKomentar()
		case 3:
			subBox("Statistik Sentimen")
			statistikSentimen()
		case 4:
			subBox("Ubah Komentar")
			ubahKomentar()
		case 5:
			subBox("Hapus Komentar")
			hapusKomentar()
		case 6:
			subBox("Cari Komentar (Sequential Search)")
			var kata string
			fmt.Print("Masukkan kata kunci: ")
			fmt.Scan(&kata)
			sequentialSearch(kata)
		case 7:
			subBox("Cari Komentar (Binary Search)")
			var kata string
			fmt.Print("Masukkan kata kunci: ")
			fmt.Scan(&kata)
			binarySearch(kata)
		case 8:
			subBox("Urutkan Komentar (Selection Sort)")
			selectionSortKomentar()
			fmt.Println("Komentar telah diurutkan berdasarkan teks.")
		case 9:
			clear()
			fmt.Println("  ____    ___     ___    ____      ____   __   __  _____   _")
			fmt.Println(" / ___|  / _ \\   / _ \\  |  _ \\    | __ )  \\ \\ / / | ____| | |")
			fmt.Println("| |  _  | | | | | | | | | | | |   |  _ \\   \\ V /  |  _|   | |")
			fmt.Println("| |_| | | |_| | | |_| | | |_| |   | |_) |   | |   | |___  |_|")
			fmt.Println(" \\____|  \\___/   \\___/  |____/    |____/    |_|   |_____| (_)")

		}

		fmt.Println("\nTekan ENTER untuk kembali ke menu...")
		fmt.Scanln()
		fmt.Scanln()
	}
}

func cari(teks string, kata string) bool {
	n := len(teks)
	m := len(kata)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && teks[i+j] == kata[j] {
			j++
		}
		if j == m {
			return true
		}
	}
	return false
}
