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

// Daftar kata sentimen
var komentarPositif = [NMAX]string{"baik", "bagus", "mantap", "hebat", "luar biasa", "puas", "senang", "terbaik", "positif", "wow", "suka", "menakjubkan", "inspiratif", "bermanfaat"}
var komentarNegatif = [NMAX]string{"jelek", "buruk", "parah", "bodoh", "jijik", "minus", "negatif", "bohong", "menyesal", "gila", "tidak puas", "jelek banget", "benci", "tidak bermanfaat", "mati"}

// fungsi untuk menganalisa sentimen berdasarkan teks komentar
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

// fungsi untuk menginput komentar
func masukKomentar() {
	var usn, teks string
	for {

		fmt.Print("\nMasukkan ID (atau # untuk keluar): ")
		fmt.Scan(&usn)
		if usn == "#" {
			return
		}
		fmt.Print("Masukkan Komentar (atau # untuk keluar): ")
		fmt.Scan(&teks)
		if teks == "#" {
			return
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
			return
		}
	}
}

// Menampilkan komentar
func tampilKomentar() {
	var i int
	fmt.Println("\nDaftar Komentar:")
	for i = 0; i < jumlahKomentar; i++ {
		fmt.Printf("[%s] %s => Sentimen: %s\n", daftarKomentar[i].usn, daftarKomentar[i].teks, daftarKomentar[i].sentimen)
	}
}

// Fungsi untuk mencari kata dalam teks
func cari(teks string, kata string) bool {
	var n, m, i, j int

	n = len(teks)
	m = len(kata)

	for i = 0; i <= n-m; i++ {
		j = 0
		for j < m && teks[i+j] == kata[j] {
			j++
		}
		if j == m {
			return true
		}
	}
	return false
}

// Menampilkan statistik sentimen
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

// Mengubah komentar berdasarkan ID
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
	fmt.Println("Komentar tidak ditemukan.")
}

// Menghapus komentar berdasarkan ID
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

// Menurutkan komentar menggunakan metode selection sort
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

// Mengurutkan komentar menggunakan metode insertion sort
func insertionSortKomentar() {
	var i, j int
	var temp komentar

	for i = 1; i < jumlahKomentar; i++ {
		temp = daftarKomentar[i]
		j = i - 1

		// Geser elemen yang lebih besar dari temp.teks ke kanan
		for j >= 0 && daftarKomentar[j].teks > temp.teks {
			daftarKomentar[j+1] = daftarKomentar[j]
			j--
		}

		// Tempatkan temp pada posisi yang tepat
		daftarKomentar[j+1] = temp
	}
}
// Mencari komentar menggunakan metode sequential search
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

// Mencari komentar menggunakan metode binary search
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
			return
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

// Membersihkan layar
func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Mengubah teks menjadi warna pink
func colorPink(text string) string {
	return "\033[38;2;209;71;124m" + text + "\033[0m"
}

// Menampilkan menu utama
func box() {
	clear()
	fmt.Println(colorPink("╔══════════════════════════════════════════════════════════════╗"))
	fmt.Println(colorPink("║                   Aplikasi Analisis Sentimen                 ║"))
	fmt.Println(colorPink("║                  Created by : Fauzan & Mario                 ║"))
	fmt.Println(colorPink("╚══════════════════════════════════════════════════════════════╝"))

	fmt.Println(colorPink("╔══════════════════════════════════════════════════════════════╗"))
	fmt.Println(colorPink("║                                                              ║"))
	fmt.Println(colorPink("║                           MAIN MENU                          ║"))
	fmt.Println(colorPink("║                                                              ║"))
	fmt.Println(colorPink("╠══════════════════════════════════════════════════════════════╣"))
	fmt.Println(colorPink("║     Selamat datang di aplikasi analisis komentar online.     ║"))
	fmt.Println(colorPink("║                                                              ║"))
	fmt.Println(colorPink("║     1. Masukkan Komentar                                     ║"))
	fmt.Println(colorPink("║     2. Tampilkan Komentar                                    ║"))
	fmt.Println(colorPink("║     3. Statistik Sentimen                                    ║"))
	fmt.Println(colorPink("║     4. Ubah Komentar                                         ║"))
	fmt.Println(colorPink("║     5. Hapus Komentar                                        ║"))
	fmt.Println(colorPink("║     6. Cari Komentar (Sequential Search)                     ║"))
	fmt.Println(colorPink("║     7. Cari Komentar (Binary Search)                         ║"))
	fmt.Println(colorPink("║     8. Urutkan Komentar (Selection Sort berdasarkan teks)    ║"))
	fmt.Println(colorPink("║     9. Urutkan Komentar (Insertion Sort berdasarkan teks)    ║"))                                        
	fmt.Println(colorPink("║     10. Keluar                                               ║"))
	fmt.Println(colorPink("║                                                              ║"))
	fmt.Println(colorPink("║     Silakan pilih menu yang Anda butuhkan.                   ║"))
	fmt.Println(colorPink("╚══════════════════════════════════════════════════════════════╝"))
	fmt.Print(colorPink("SELECT> "))
}

// Menampilkan subjudul dengan border
func subBox(judul string) {
	clear()
	fmt.Println(colorPink("╔══════════════════════════════════════════════════════════════╗"))
	fmt.Printf("║ %-60s ║\n", judul)
	fmt.Println(colorPink("╚══════════════════════════════════════════════════════════════╝"))
}

// program utama
func main() {
	var pilih int
	for pilih != 10 {
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
			subBox("Urutkan Komentar (Insertion Sort)")
			insertionSortKomentar()
			fmt.Println("Komentar telah diurutkan menggunakan insertion sort berdasarkan teks.")
		case 10:
			clear()
			fmt.Println(colorPink("  ____    ___     ___    ____      ____   __   __  _____   _"))
			fmt.Println(colorPink(" / ___|  / _ \\   / _ \\  |  _ \\    | __ )  \\ \\ / / | ____| | |"))
			fmt.Println(colorPink("| |  _  | | | | | | | | | | | |   |  _ \\   \\ V /  |  _|   | |"))
			fmt.Println(colorPink("| |_| | | |_| | | |_| | | |_| |   | |_) |   | |   | |___  |_|"))
			fmt.Println(colorPink(" \\____|  \\___/   \\___/  |____/    |____/    |_|   |_____| (_)"))

		}

		fmt.Println("\nTekan ENTER untuk kembali ke menu...")
		fmt.Scanln()
		fmt.Scanln()
	}
}
