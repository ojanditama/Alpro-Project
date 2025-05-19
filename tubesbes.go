package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const NMAX int = 1000

type question struct {
	id               int
	tag              string
	pertanyaan       string
	tanggapan        []string
	date             string
	author           string
	jumlah_tanggapan int
}

type dataAccount struct {
	nama        string
	username    string
	password    string
	PIN         int
	jumlah_post int
}

type acc [NMAX]dataAccount
type que [NMAX]question

func main() {
	var account acc
	var question que
	var totalAccount, totalPertanyaan int
	header(account, totalAccount, totalPertanyaan, question)
}

func header(A acc, T int, P int, Q que) {
	list := []string{"1. Pasien", "2. Dokter", "3. Lihat Forum Konsultasi", "4. Load Backup", "5. Exit"}
	pesan := []string{"Silahkan pilih keperluan anda."}
	menu("Selamat datang di aplikasi konsultasi online.", list, pesan, "MAIN", false, &A, T, P, &Q)
	main_menu(&A, &T, &P, &Q)
}

// Procedure main menu dari program aplikasi
func main_menu(akun *acc, T *int, P *int, question *que) {
	var pilihan int
	fmt.Scanln(&pilihan)
	pesan := []string{"Silahkan pilih keperluan anda."}
	switch pilihan {
	case 1:
		list := []string{"1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back"}
		menu("Selamat datang di menu pasien.", list, pesan, "PASIEN", false, akun, *T, *P, question)
		pasien(*akun, T, P, *question)
	case 2:
		list := []string{""}
		pesan := []string{"Perkenalkan diri anda sebagai dokter."}
		menu("Selamat datang di menu dokter.", list, pesan, "DOKTER", true, akun, *T, *P, question)
		dokter(*akun, T, P, *question)
	case 3:
		PrintForum(*akun, *T, *P, *question, "PENGGUNA", "", false)
	case 4:
		load(akun, T, P, question)
		list := []string{"1. Pasien", "2. Dokter", "3. Lihat Forum Konsultasi", "4. Load Backup", "5. Exit"}
		menu("Selamat datang di aplikasi konsultasi online.", list, pesan, "MAIN", false, akun, *T, *P, question)
		main_menu(akun, T, P, question)
	case 5:
		bye()
	default:
		list := []string{"1. Pasien", "2. Dokter", "3. Lihat Forum Konsultasi", "4. Load Backup", "5. Exit"}
		menu("Masukkan angka sesuai option.", list, pesan, "MAIN", false, akun, *T, *P, question)
		main_menu(akun, T, P, question)
	}
}

// Procedure untuk pilihan pengguna apakah pengguna ingin login akun, buat akun, atau recovery akun
func pasien(akun acc, T *int, P *int, question que) {
	var pilihan int
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		list := []string{"Login Akun Pasien"}
		pesan := []string{"Silahkan masukkan username dan password anda."}
		menu("", list, pesan, "PASIEN", true, &akun, *T, *P, &question)
		loginAccount(akun, *T, *P, question)
	case 2:
		list := []string{"Buat Akun Baru Pasien"}
		pesan := []string{"Silahkan masukkan username dan password anda."}
		menu("", list, pesan, "PASIEN", true, &akun, *T, *P, &question)
		makeAcc(&akun, T, P, &question)
	case 3:
		list := []string{"Recovery Akun Pasien."}
		pesan := []string{"Silahkan masukkan username dan password anda."}
		menu("", list, pesan, "PASIEN", true, &akun, *T, *P, &question)
		recovery(&akun, *T, *P, &question)
	case 4:
		list := []string{"1. Pasien", "2. Dokter", "3. Lihat Forum Konsultasi", "4. Load Backup", "5. Exit"}
		pesan := []string{"Silahkan pilih kerpeluan anda"}
		menu("Selamat datang di aplikasi konsultasi online.", list, pesan, "MAIN", false, &akun, *T, *P, &question)
		main_menu(&akun, T, P, &question)
	default:
		list := []string{"1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back"}
		pesan := []string{"Silahkan pilih kerpeluan anda"}
		menu("Masukkan angka sesuai option.", list, pesan, "PASIEN", false, &akun, *T, *P, &question)
		pasien(akun, T, P, question)
	}
}

// Procedure untuk perkenalan nama dokter sebagai panggilan
func dokter(akun acc, T *int, P *int, question que) {
	var nama string

	fmt.Printf("%24s %s", "", "Masukkan nama Anda sebagai dokter : ")
	nama = scantext()
	welcome_text := "Selamat datang dr." + nama + ", di menu dokter : "
	list := []string{"1. Lihat Forum Konsultasi", "2. Cari Pasien", "3. Kembali"}
	pesan := []string{"Silahkan pilih kerpeluan anda"}
	menu(welcome_text, list, pesan, "DOKTER", false, &akun, *T, *P, &question)
	main_dokter(akun, T, P, question, nama)
}

// Procedure menu utama untuk dokter
func main_dokter(akun acc, T *int, P *int, question que, nama string) {
	var pilihan int
	fmt.Scanln(&pilihan)
	pesan := []string{"Silahkan pilih kerpeluan anda"}
	switch pilihan {
	case 1:
		PrintForum(akun, *T, *P, question, "DOKTER", nama, false)
	case 2:
		if *T == 0 {
			pesan := []string{"1. Kembali"}
			list := []string{"Tidak ada pasien yang tersedia saat ini."}
			menu("", list, pesan, "DOKTER", false, &akun, *T, *P, &question)
		} else {
			SelectionSortAccount(&akun, *T)
			pesan := []string{"Silahkan masukkan nama pasien."}
			list := []string{"List nama pasien : "}
			for i := 0; i < *T; i++ {
				list = append(list, fmt.Sprintf("%d. %s", i+1, akun[i].nama))
			}
			menu("Mencari data pasien sesuai dengan nama pasien.", list, pesan, "DOKTER", true, &akun, *T, *P, &question)
			idx := caripasien(&akun, *T, nama)
			if idx != -1 {
				if akun[idx].jumlah_post == 0 {
					kata1 := fmt.Sprintf("Pasien atas nama %s dengan username %s", akun[idx].nama, akun[idx].username)
					list := []string{kata1, "tidak pernah melakukan posting pertanyaan."}
					pesan := []string{"1. Kembali"}
					menu("Nama pasien telah di temukan.", list, pesan, "DOKTER", false, &akun, *T, *P, &question)
				} else {
					kata1 := fmt.Sprintf("Pasien atas nama %s dengan username %s", akun[idx].nama, akun[idx].username)
					kata2 := fmt.Sprintf("telah memposting sebanyak %d pertanyaan.", akun[idx].jumlah_post)
					kata3 := fmt.Sprintf("List pertanyaan oleh %s :", akun[idx].nama)
					header := fmt.Sprintf("Tag%-5sLast Post%-14sAuthor%-20sTanggapan", "", "", "")
					list := []string{kata1, kata2, " ", kata3, " ", header}
					for i := 0; i < *P; i++ {
						if strings.EqualFold(question[i].author, akun[idx].nama) {
							kata4 := fmt.Sprintf("%-7s %-23s%-26s%d", question[i].tag, question[i].date, question[i].author, question[i].jumlah_tanggapan)
							list = append(list, kata4)
						}
					}
					pesan := []string{"1. Kembali"}
					menu("Nama pasien telah di temukan.", list, pesan, "DOKTER", false, &akun, *T, *P, &question)
				}
			} else {
				list := []string{}
				pesan := []string{"1. Kembali"}
				menu("Nama pasien tidak di temukan.", list, pesan, "DOKTER", false, &akun, *T, *P, &question)
			}
		}
		fmt.Scanln(&pilihan)
		welcome_text := "Selamat datang dr." + nama + ", di menu dokter : "
		list := []string{"1. Lihat Forum Konsultasi", "2. Cari Pasien", "3. Kembali"}
		pesan := []string{"Silahkan pilih kerpeluan anda"}
		menu(welcome_text, list, pesan, "DOKTER", false, &akun, *T, *P, &question)
		main_dokter(akun, T, P, question, nama)
	case 3:
		list := []string{"1. Pasien", "2. Dokter", "3. Lihat Forum Konsultasi", "4. Load Backup", "5. Exit"}
		menu("Selamat datang di aplikasi konsultasi online.", list, pesan, "MAIN", false, &akun, *T, *P, &question)
		main_menu(&akun, T, P, &question)
	default:
		list := []string{"1. Lihat Forum Konsultasi", "2. Cari Pasien", "3. Kembali"}
		menu("Masukkan angka sesuai option.", list, pesan, "DOKTER", false, &akun, *T, *P, &question)
		main_dokter(akun, T, P, question, nama)
	}
}

// Menu utama dari segalanya, dengan cara menggunakan procedure ini maka kita tidak perlu membuat menu yang sama lagi berulang ulang kali
// Hanya memanggil procedure dengan parameter yang diminta, maka akan menampilkan menu yang diinginkan
func menu(header string, kata []string, message []string, menuType string, hide bool, akun *acc, T int, P int, question *que) {
	clear()
	color("a")
	title("TUGAS BESAR ALGORITMA PEMPROGRAMAN")
	fmt.Printf("%25s╔═════════════════════════════════════════════════════════════════════════╗\n", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s║%53s%-20s║\n", "", "Aplikasi Konsultasi Kesehatan", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s║%55s %17s║\n", "", "Created by : • Faisal Ihsan Santoso", "")
	fmt.Printf("%25s║%59s %13s║\n", "", "• Arie Farchan Fyrzatullah", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	fmt.Printf("%25s╠═════════════════════════════════════════════════════════════════════════╣\n", "")
	fmt.Printf("%25s║%73s║\n", "", "")
	if header != "" {
		fmt.Printf("%25s║  %-71s║\n", "", header)
		fmt.Printf("%25s║%73s║\n", "", "")
	}
	for _, word := range kata {
		if word != "" {
			fmt.Printf("%25s║%4s%-69s║\n", "", "", word)
		}
	}
	fmt.Printf("%25s║%73s║\n", "", "")
	for _, pesan := range message {
		if pesan != "" {
			fmt.Printf("%25s║  %-71s║\n", "", pesan)
		}
	}
	fmt.Printf("%25s║%73s║\n", "", "")
	if menuType == "PASIEN" {
		fmt.Printf("%25s║%59s╔═════════════╣\n", "", "")
		fmt.Printf("%25s║%73s║\n", "", "║ MENU PASIEN ")
		fmt.Printf("%25s╚═══════════════════════════════════════════════════════════╩═════════════╝\n", "")
	} else if menuType == "DOKTER" {
		fmt.Printf("%25s║%59s╔═════════════╣\n", "", "")
		fmt.Printf("%25s║%73s║\n", "", "║ MENU DOKTER ")
		fmt.Printf("%25s╚═══════════════════════════════════════════════════════════╩═════════════╝\n", "")
	} else {
		fmt.Printf("%25s║%61s╔═══════════╣\n", "", "")
		fmt.Printf("%25s║%73s║\n", "", "║ MAIN MENU ")
		fmt.Printf("%25s╚═════════════════════════════════════════════════════════════╩═══════════╝\n", "")
	}
	if !hide {
		fmt.Printf("%25sSELECT> ", "")
	}
}

// Procedure utama untuk menu pasien
func main_pasien(A acc, T int, P int, username string, question que) {
	var pilihan int
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		list := []string{"Posting keluhan anda yang ingin anda tanyakan pada dokter."}
		pesan := []string{"Silahkan masukkan Judul Pertanyaan dan Pertanyaan anda."}
		menu("", list, pesan, "PASIEN", true, &A, T, P, &question)
		PostQuestion(&A, T, &P, username, &question)
	case 2:
		PrintForum(A, T, P, question, "PASIEN", username, false)
	case 3:
		list := []string{"1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back"}
		pesan := []string{"Silahkan pilih keperluan anda."}
		menu("Selamat datang di menu pasien.", list, pesan, "PASIEN", false, &A, T, P, &question)
		pasien(A, &T, &P, question)
	default:
		list := []string{"1. Konsultasi", "2. Lihat Forum Konsultasi", "3. Logout"}
		pesan := []string{"Silahkan pilih kerpeluan anda"}
		menu("Masukkan angka sesuai option.", list, pesan, "PASIEN", false, &A, T, P, &question)
		main_pasien(A, T, P, username, question)
	}
}

// Procedure untuk memposting pertanyaan dari pasien
func PostQuestion(A *acc, T int, P *int, username string, Q *que) {
	Index := findIndexUser(*A, T, username)
	var judul, pertanyaan string
	fmt.Printf("%24s %s", "", "Masukkan Judul Pertanyaan anda : ")
	fmt.Scanln(&judul)
	fmt.Printf("%24s %s", "", "Masukkan Pertanyaan Anda : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pertanyaan = scanner.Text()
	Q[*P].tag = judul
	Q[*P].pertanyaan = pertanyaan
	A[Index].jumlah_post += 1
	currentTime := time.Now()
	Q[*P].date = string(currentTime.Format("2 January 2006 15:04"))
	Q[*P].author = A[Index].nama
	Q[*P].id = *P + 1
	*P += 1
	list := []string{"1. Konsultasi", "2. Lihat Forum Konsultasi", "3. Logout"}
	pesan := []string{"Silahkan pilih kerpeluan anda"}
	menu("Pertanyaan anda berhasil di posting!", list, pesan, "PASIEN", false, A, T, *P, Q)
	main_pasien(*A, T, *P, username, *Q)
}

// Procedure untuk login akun pasien
func loginAccount(A acc, T int, P int, question que) {
	var username, password string
	var indexUser int
	fmt.Printf("%24s %s", "", "Masukkan username Anda : ")
	fmt.Scanln(&username)
	fmt.Printf("%24s %s", "", "Masukkan password Anda : ")
	fmt.Scanln(&password)

	if exist(A, T, username) {
		indexUser = findIndexUser(A, T, username)
		if A[indexUser].username == username && A[indexUser].password == password {
			welcome_text := "Selamat datang " + A[indexUser].nama + ", di menu pasien."
			list := []string{"1. Konsultasi", "2. Lihat Forum Konsultasi", "3. Logout"}
			pesan := []string{"Silahkan pilih kerpeluan anda"}
			menu(welcome_text, list, pesan, "PASIEN", false, &A, T, P, &question)
			main_pasien(A, T, P, username, question)
			//Login Berhasil
		} else {
			list := []string{"1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back"}
			pesan := []string{"Silahkan pilih kerpeluan anda"}
			menu("Password anda salah, gunakan PIN untuk recovery akun anda.", list, pesan, "PASIEN", false, &A, T, P, &question)
			pasien(A, &T, &P, question)
			//Login Gagal
		}
	} else {
		list := []string{"1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back"}
		pesan := []string{"Silahkan pilih kerpeluan anda"}
		menu("Akun anda tidak ditemukan, silahkan buat akun baru.", list, pesan, "PASIEN", false, &A, T, P, &question)
		pasien(A, &T, &P, question)
		//Akun tidak ditemukan
	}
}

// Procedure untuk membuat akun
func makeAcc(A *acc, T *int, P *int, question *que) {
	var nama, username, password string
	var pin int
	fmt.Printf("%24s %s", "", "Masukkan nama Anda : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama = scanner.Text()
	fmt.Printf("%24s %s", "", "Masukkan username Anda : ")
	fmt.Scanln(&username)
	fmt.Printf("%24s %s", "", "Masukkan password Anda : ")
	fmt.Scanln(&password)
	fmt.Printf("%24s %s", "", "Masukkan PIN Anda (4 Digit) : ")
	fmt.Scanln(&pin)

	if exist(*A, *T, username) { //Buat akun gagal username sudah di gunakan
		list := []string{"Buat Akun Baru Pasien"}
		pesan := []string{"Silahkan masukkan username dan password anda."}
		menu("Username telah di gunakan, silahkan gunakan username lain.", list, pesan, "PASIEN", true, A, *T, *P, question)
		makeAcc(A, T, P, question)
	} else {
		A[*T].nama = nama
		A[*T].username = username
		A[*T].password = password
		A[*T].PIN = pin
		*T += 1
		list := []string{"1. Login Akun Pasien", "2. Buat Akun Baru Pasien", "3. Recovery Akun Pasien", "4. Back"}
		pesan := []string{"Silahkan pilih kerpeluan anda"}
		menu("Akun berhasil dibuat, Silahkan login menggunakan akun anda!", list, pesan, "PASIEN", false, A, *T, *P, question)
		pasien(*A, T, P, *question)
	}
}

// Sequential Search untuk mencari apakah username x ada di data array struct atau tidak
func exist(A acc, T int, usr string) bool {
	for i := 0; i < T; i++ {
		if usr == A[i].username {
			return true
		}
	}
	return false
}

// Option untuk pasien yang lupa password, hanya menggunakan username dan PIN maka pasien bisa mereset password
func recovery(A *acc, T int, P int, question *que) {
	var username string
	var passwordBaru string
	var indexUser, pin int
	fmt.Printf("%24s %s", "", "Masukkan username Anda : ")
	fmt.Scanln(&username)
	list := []string{"1. Login Akun Pasien", "2. Buat Akun baru Pasien", "3. Recovery Akun Pasien", "4. Back"}
	pesan := []string{"Silahkan pilih keperluan anda."}

	if exist(*A, T, username) {
		indexUser = findIndexUser(*A, T, username)
		fmt.Printf("%24s %s", "", "Masukkan PIN Anda : ")
		fmt.Scanln(&pin)
		if A[indexUser].PIN == pin {
			fmt.Printf("%24s %s", "", "Masukkan password baru Anda : ")
			fmt.Scanln(&passwordBaru)
			A[indexUser].password = passwordBaru
			menu("Akun anda berhasil di Recovery, silahkan login dengan password baru.", list, pesan, "PASIEN", false, A, T, P, question)
			pasien(*A, &T, &P, *question)
		} else {
			menu("Anda memasukan PIN yang salah.", list, pesan, "PASIEN", false, A, T, P, question)
			pasien(*A, &T, &P, *question)
		}
	} else {
		menu("Akun anda tidak ditemukan, silahkan buat akun baru.", list, pesan, "PASIEN", false, A, T, P, question)
		pasien(*A, &T, &P, *question)
	}
}

// Mirip dengan fmt.Scan() tetapi ini menggunakan library bufio sehingga bisa mengambil input lebih dari satu kata
// Seperti contoh input Faisal Ihsan kalau biasa hanyak di ambil Faisal, tetapi dengan bufio kita bisa mengambil Faisal Ihsan
func scantext() string {
	var nama, text string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama = scanner.Text()
	text = nama

	return text
}

// Menampilkan menu forum sesuai dengan usertype "PENGGUNA" / "PASIEN" / "DOKTER"
func PrintForum(A acc, T int, P int, Q que, usertype, username string, NotFound bool) {
	var pilihan int
	if P == 0 {
		list := []string{"Forum pertanyaan masih kosong."}
		pesan := []string{"1. Kembali"}
		menu("Selamat datang di Forum Konsultasi.", list, pesan, "MAIN", false, &A, T, P, &Q)
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			if usertype == "PENGGUNA" {
				header(A, T, P, Q)
			} else if usertype == "PASIEN" {
				indexUser := findIndexUser(A, T, username)
				welcome_text := "Selamat datang " + A[indexUser].nama + ", di menu pasien."
				list := []string{"1. Konsultasi", "2. Lihat Forum Konsultasi", "3. Logout"}
				pesan := []string{"Silahkan pilih kerpeluan anda"}
				menu(welcome_text, list, pesan, "PASIEN", false, &A, T, P, &Q)
				main_pasien(A, T, P, username, Q)
			} else {
				welcome_text := "Selamat datang dr." + username + ", di menu dokter : "
				list := []string{"1. Lihat Forum Konsultasi", "2. Cari Pasien", "3. Kembali"}
				pesan := []string{"Silahkan pilih kerpeluan anda"}
				menu(welcome_text, list, pesan, "DOKTER", false, &A, T, P, &Q)
				main_dokter(A, &T, &P, Q, username)
			}
		default:
			list := []string{"Forum pertanyaan masih kosong."}
			pesan := []string{"1. Kembali"}
			menu("Masukkan angka sesuai option.", list, pesan, "MAIN", false, &A, T, P, &Q)
			PrintForum(A, T, P, Q, usertype, username, false)
		}
	} else {
		kata := fmt.Sprintf("Tag%-8sLast Post%-14sAuthor%-19sTanggapan", "", "", "")
		list := []string{kata}
		for i := 0; i < P; i++ {
			if i >= 9 {
				kata2 := fmt.Sprintf("%d. %-6s %-23s%-26s%d", i+1, Q[i].tag, Q[i].date, Q[i].author, Q[i].jumlah_tanggapan)
				list = append(list, kata2)
			} else {
				kata2 := fmt.Sprintf("%d. %-7s %-23s%-26s%d", i+1, Q[i].tag, Q[i].date, Q[i].author, Q[i].jumlah_tanggapan)
				list = append(list, kata2)
			}
		}
		pesan := []string{" ", "Silahkan pilih keperluan anda.", " ", "1. Lihat Pertanyaan", "2. Cari Pertanyaan dengan Tag", "3. Urutkan dari tanggapan terbanyak", "4. Urutkan dari tanggapan sedikit", "5. Kembali"}
		if NotFound {
			menu("Pertanyaan tidak di temukan.", list, pesan, "MAIN", false, &A, T, P, &Q)
		} else {
			if usertype == "PENGGUNA" {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "MAIN", false, &A, T, P, &Q)
			} else if usertype == "DOKTER" {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "DOKTER", false, &A, T, P, &Q)
			} else {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "PASIEN", false, &A, T, P, &Q)
			}
		}
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			var id int
			pesan := []string{"Pilih pertanyaan yang anda ingin liat."}
			if usertype == "PENGGUNA" {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "MAIN", true, &A, T, P, &Q)
			} else if usertype == "DOKTER" {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "DOKTER", true, &A, T, P, &Q)
			} else {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "PASIEN", true, &A, T, P, &Q)
			}
			fmt.Printf("%24s %s", "", "Masukkan nomor Pertanyaan yang ingin Anda lihat : ")
			fmt.Scanln(&id)
			ViewQuestion(A, T, P, id, &Q, usertype, username)
		case 2:
			var tag string
			pesan := []string{"Berikan TAG pertanyaan yang anda ingin liat."}
			if usertype == "PENGGUNA" {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "MAIN", true, &A, T, P, &Q)
			} else if usertype == "DOKTER" {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "DOKTER", true, &A, T, P, &Q)
			} else {
				menu("Selamat datang di Forum Konsultasi.", list, pesan, "PASIEN", true, &A, T, P, &Q)
			}
			fmt.Printf("%24s %s", "", "Masukkan TAG Pertanyaan yang ingin Anda lihat : ")
			fmt.Scanln(&tag)
			id := findIndexQuestion2(Q, P, tag)
			ViewQuestion(A, T, P, id+1, &Q, usertype, username)
		case 3:
			InsertionSortQuestion(&Q, P) // Mengunakan metode Insertion Sort untuk mengurutan dari tanggapan terbanyak
			PrintForum(A, T, P, Q, usertype, username, false)
		case 4:
			SelectionSortQuestion(&Q, P) // Mengunakan metode Selection Sort untuk mengurutan dari tanggapan tersedikit
			PrintForum(A, T, P, Q, usertype, username, false)
		case 5:
			if usertype == "PENGGUNA" {
				header(A, T, P, Q)
			} else if usertype == "PASIEN" {
				indexUser := findIndexUser(A, T, username)
				welcome_text := "Selamat datang " + A[indexUser].nama + ", di menu pasien."
				list := []string{"1. Konsultasi", "2. Lihat Forum Konsultasi", "3. Logout"}
				pesan := []string{"Silahkan pilih kerpeluan anda"}
				menu(welcome_text, list, pesan, "PASIEN", false, &A, T, P, &Q)
				main_pasien(A, T, P, username, Q)
			} else {
				welcome_text := "Selamat datang dr." + username + ", di menu dokter : "
				list := []string{"1. Lihat Forum Konsultasi", "2. Cari Pasien", "3. Kembali"}
				pesan := []string{"Silahkan pilih kerpeluan anda"}
				menu(welcome_text, list, pesan, "DOKTER", false, &A, T, P, &Q)
				main_dokter(A, &T, &P, Q, username)
			}
		default:
			PrintForum(A, T, P, Q, usertype, username, false)
		}
	}
}

// Tampilkan pertanyaan serta tanggapan dokter/pasien
func ViewQuestion(A acc, T int, P int, id int, Q *que, usertype, username string) {
	var pilihan int
	indexQ := findIndexQuestion(*Q, P, id)
	if indexQ == -1 {
		PrintForum(A, T, P, *Q, usertype, username, true)
	}
	var list []string
	if Q[indexQ].jumlah_tanggapan == 0 {
		list = []string{"Belum ada tanggapan untuk pertanyaan ini."}
		if usertype == "PENGGUNA" {
			pesan := []string{"1. Kembali"}
			menu(Q[indexQ].pertanyaan, list, pesan, "MAIN", false, &A, T, P, Q)
		} else {
			pesan := []string{"1. Beri tanggapan", "2. Kembali"}
			menu(Q[indexQ].pertanyaan, list, pesan, "MAIN", false, &A, T, P, Q)
		}
	} else {
		kata := fmt.Sprintf("%-27sTANGGAPAN", "")
		list = []string{" ", kata, " "}
		for i := 0; i < Q[indexQ].jumlah_tanggapan; i++ {
			kata2 := fmt.Sprintf(Q[indexQ].tanggapan[i])
			list = append(list, kata2)
		}
		if usertype == "PENGGUNA" {
			pesan := []string{"1. Kembali"}
			menu(Q[indexQ].pertanyaan, list, pesan, "MAIN", false, &A, T, P, Q)
		} else {
			pesan := []string{"1. Beri tanggapan", "2. Kembali"}
			menu(Q[indexQ].pertanyaan, list, pesan, "MAIN", false, &A, T, P, Q)
		}
	}
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		if usertype == "PENGGUNA" {
			PrintForum(A, T, P, *Q, usertype, username, false)
		} else {
			pesan := []string{"Berikan tanggapan anda mengenai pertanyaan di atas."}
			menu(Q[indexQ].pertanyaan, list, pesan, "MAIN", true, &A, T, P, Q)
			fmt.Printf("%24s %s", "", "Masukkan tanggapan anda : ")
			text := scantext()
			if usertype == "PASIEN" {
				idx := findIndexUser(A, T, username)
				text = A[idx].nama + " : " + text
			} else {
				text = "dr." + username + " : " + text
			}
			Q[indexQ].tanggapan = append(Q[indexQ].tanggapan, text)
			Q[indexQ].jumlah_tanggapan++
			currentTime := time.Now()
			Q[indexQ].date = string(currentTime.Format("2 January 2006 15:04"))
			ViewQuestion(A, T, P, id, Q, usertype, username)
		}
	case 2:
		PrintForum(A, T, P, *Q, usertype, username, false)
	default:
		ViewQuestion(A, T, P, id, Q, usertype, username)
	}
}

// Insertion Sort urutan dari tanggapan terbanyak
func InsertionSortQuestion(Q *que, P int) {
	for i := 1; i < P; i++ {
		num := Q[i]
		j := i - 1
		for j >= 0 && Q[j].jumlah_tanggapan < num.jumlah_tanggapan {
			Q[j+1] = Q[j]
			j = j - 1
		}
		Q[j+1] = num
	}
	for i := 0; i < P; i++ {
		Q[i].id = i + 1
	}
}

// SelectionSort urutan dari tanggapan tersedikit
func SelectionSortQuestion(Q *que, P int) {
	for i := 0; i < P; i++ {
		idx := i
		for x := i; x < P; x++ {
			if Q[idx].jumlah_tanggapan > Q[x].jumlah_tanggapan {
				idx = x
			}
		}
		Q[i].id, Q[idx].id = Q[idx].id, Q[i].id
		Q[i], Q[idx] = Q[idx], Q[i]
	}
}

// SelectionSort urutan nama pasien sesuai abjad
func SelectionSortAccount(A *acc, T int) {
	for i := 0; i < T; i++ {
		min := i
		for x := i; x < T; x++ {
			if A[min].nama > A[x].nama {
				min = x
			}
		}
		A[i], A[min] = A[min], A[i]
	}
}

// Binary search untuk mencari nama pasien yang sudah di urutkan sesuai abjad
func caripasien(A *acc, T int, nama string) int {
	kiri := 0
	kanan := T - 1
	var nama_s string
	tengah := (kiri + kanan) / 2
	nama_idx := strings.ToLower(A[tengah].nama)
	nama_cari := strings.ToLower(nama_s)
	fmt.Printf("%24s %s", "", "Masukkan Nama Lengkap Pasien yang ingin anda cari : ")
	nama_s = scantext()
	for kiri <= kanan && nama_cari != nama_idx {
		nama_idx = strings.ToLower(A[tengah].nama)
		nama_cari = strings.ToLower(nama_s)
		if nama_cari < nama_idx {
			kanan = tengah - 1
		} else if nama_cari > nama_idx {
			kiri = tengah + 1
		} else {
			return tengah
		}
		tengah = (kiri + kanan) / 2
	}
	return -1
}

// Sequential Search untuk mencari index dari suatu username
func findIndexUser(A acc, T int, usr string) int {
	for i := 0; i < T; i++ {
		if usr == A[i].username {
			return i
		}
	}
	return 0
}

// Sequential Search untuk mencari index dari suatu pertanyaan dengan id
func findIndexQuestion(Q que, P, id int) int {
	for i := 0; i < P; i++ {
		if id == Q[i].id {
			return i
		}
	}
	return -1
}

// Sequential Search untuk mencari index dari suatu pertanyaan dengan tag
func findIndexQuestion2(Q que, P int, tag string) int {
	for i := 0; i < P; i++ {
		fmt.Println(Q[i].tag, tag)
		if strings.EqualFold(Q[i].tag, tag) {
			return i
		}
	}
	return -1
}

// Procedure untuk membersihkan console program
func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Procedure untuk warna console program
func color(text string) {
	cmd := exec.Command("cmd", "/c", "color ", text)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Procedure untuk judul console program
func title(text string) {
	cmd := exec.Command("cmd", "/c", fmt.Sprintf("title %s", text))
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Procedure untuk keluar program
func bye() {
	clear()
	fmt.Println("  ____    ___     ___    ____      ____   __   __  _____   _")
	fmt.Println(" / ___|  / _ \\   / _ \\  |  _ \\    | __ )  \\ \\ / / | ____| | |")
	fmt.Println("| |  _  | | | | | | | | | | | |   |  _ \\   \\ V /  |  _|   | |")
	fmt.Println("| |_| | | |_| | | |_| | | |_| |   | |_) |   | |   | |___  |_|")
	fmt.Println(" \\____|  \\___/   \\___/  |____/    |____/    |_|   |_____| (_)")
}

// Procedure untuk mengload data struct array, sehingga tidak membuang waktu ketika mempraktekan
func load_account(name, username, password string, pin, jumlah int) dataAccount {
	return dataAccount{nama: name, username: username, password: password, PIN: pin, jumlah_post: jumlah}
}

// Procedure untuk mengload data struct array, sehingga tidak membuang waktu ketika mempraktekan
func load_question(id int, tag, pertanyaan string, tanggapan []string, date, author string) question {
	return question{id: id, tag: tag, pertanyaan: pertanyaan, tanggapan: tanggapan, date: date, author: author}
}

// Procedure untuk mengload data struct array, sehingga tidak membuang waktu ketika mempraktekan
func load(A *acc, T *int, P *int, Q *que) {
	//urutan : NAMA, USERNAME, PASSWORD, PIN, JUMLAH POST
	A[*T] = load_account("Jarjit Susanto", "jarjit", "jar123", 1234, 0)
	*T += 1
	A[*T] = load_account("Faisal Ihsan Santoso", "faisal", "1234", 1122, 4)
	*T += 1
	A[*T] = load_account("Arie Farchan Fyrzatullah", "arie", "ariex", 1234, 3)
	*T += 1
	A[*T] = load_account("Rizki Nata", "nata", "riznat", 9876, 2)
	*T += 1
	A[*T] = load_account("Jundi Haq", "jundHD", "junjun", 2222, 3)
	*T += 1
	A[*T] = load_account("Bayu Putra", "bayu", "bay", 4444, 1)
	*T += 1
	//urutan : ID, TAG, PERTANYAAN, TANGAPPAN (list), TANGGAL LAST POST, AUTHOR
	list := []string{"dr.budi : Anda bisa mengurangi gejala alergi musiman", "Faisal Ihsan Santoso : Terima kasih atas informasinya, dr.budi.", "dr.budi : Obat antihistamin dan dekongestan bisa membantu gejala."}
	Q[*P] = load_question(1, "Alergi", "Apa yang bisa saya lakukan untuk mengurangi gejala alergi musiman?", list, "1 Desember 2019 15:04", "Faisal Ihsan Santoso")
	Q[*P].jumlah_tanggapan = 3
	*P += 1

	list2 := []string{"dr.hadi : Jika batuk Anda disebabkan oleh alergi, hindari asap rokok.", "dr.eka : Minum banyak air, menjaga kebersihan udara di sekitar Anda.", "dr.Gita : Ada banyak jenis obat, tergantung pada penyebab batuk Anda.", "Arie Farchan Fyrzatullah : Terimakasih atas informasinya, para dokter"}
	Q[*P] = load_question(2, "Batuk", "Apakah batuk yang saya alami ini perlu diperiksakan lebih lanjut?", list2, "20 Februari 2021 11:55", "Arie Farchan Fyrzatullah")
	Q[*P].jumlah_tanggapan = 4
	*P += 1

	list3 := []string{"dr.dian : Efek samping dari antihistamin adalah mulut kering."}
	Q[*P] = load_question(3, "Obat", "Apakah ada efek samping dari obat yang sedang saya konsumsi?", list3, "9 Maret 2024 9:34", "Faisal Ihsan Santoso")
	Q[*P].jumlah_tanggapan = 1
	*P += 1

	list4 := []string{"dr.Lina : Minum banyak air dan hindari menahan kencing.", "Arie Farchan Fyrzatullah : Apakah ada makanan yang membantu?", "dr. Lina : Iya. Jus cranberry juga bisa membantu mencegah infeksi."}
	Q[*P] = load_question(4, "Anak", "Bagaimana cara menangani demam pada anak-anak?", list4, "17 Maret 2024 7:25", "Arie Farchan Fyrzatullah")
	Q[*P].jumlah_tanggapan = 3
	*P += 1

	list5 := []string{"dr.Wisnu : Ya, untuk mencegah komplikasi flu.", "Jundi Haq : Siapa yang paling perlu vaksin flu?", "dr. Wisnu : Orang tua dan anak kecil dengan kondisi medis kronis."}
	Q[*P] = load_question(5, "Vaksin", "Bagaimana cara menangani demam pada anak-anak?", list5, "27 Mei 2023 13:21", "Jundi Haq")
	Q[*P].jumlah_tanggapan = 3
	*P += 1

	list6 := []string{"dr.Udin : Minimal 150 menit per minggu.", "Rizki Nata:enis olahraga yang disarankan?", "dr. Udin : Aerobik seperti berjalan cepat atau bersepeda."}
	Q[*P] = load_question(6, "Jantung", " Seberapa sering saya harus berolahraga untuk menjaga kesehatan jantung??", list6, "2 November 2022 19:17", "Rizki Nata")
	Q[*P].jumlah_tanggapan = 3
	*P += 1

	list7 := []string{"dr.Oka : Olahraga ringan dan obat pereda nyeri.", "Faisal Ihsan Santoso : Apakah perlu fisioterapi?", "dr. Oka : Fisioterapi bisa sangat membantu"}
	Q[*P] = load_question(7, "Nyeri", "Bagaimana cara mengatasi nyeri sendi yang kronis?", list7, "5 Januari 2023 13:01", "Faisal Ihsan Santoso")
	Q[*P].jumlah_tanggapan = 3
	*P += 1

	list8 := []string{"dr.Joko : Kelelahan, pucat, dan sesak napas."}
	Q[*P] = load_question(8, "Anemia", "Apa saja gejala anemia yang harus saya perhatikan?", list8, "13 September 2023 9:41", "Bayu Putra")
	Q[*P].jumlah_tanggapan = 1
	*P += 1

	list9 := []string{"dr.Zaki : Olahraga, meditasi, dan teknik relaksasi.", "Faisal Ihsan Santoso : Terimakasih atas sarannya dr.Zaki"}
	Q[*P] = load_question(9, "Stres", "Bagaimana cara mengelola stres yang berlebihan?", list9, "30 April 2022 23:51", "Faisal Ihsan Santoso")
	Q[*P].jumlah_tanggapan = 2
	*P += 1

	list10 := []string{"dr.Vira : Jika diet Anda seimbang, tidak perlu.", "Jundi Haq : Kapan suplemen diperlukan?", "dr.Vira : Saat ada kekurangan vitamin tertentu.", "Jundi Haq : Saya dalam keadan sehat", "dr.Vira : Maka, suplemen tidak dibutuhkan", "Jundi Haq : Terimakasih Dok"}
	Q[*P] = load_question(10, "Diet", "Apakah saya perlu suplemen vitamin tambahan?", list10, "29 februari 2024 8:45", "Jundi Haq")
	Q[*P].jumlah_tanggapan = 6
	*P += 1

	list11 := []string{"dr.Andy : Coba olahraga, meditasi, dan teknik relaksasi.", "Arie Farchan Fyrzatullah : Apakah ada obat yang perlu dikonsumsi", "dr.Andy : Saya sarankan konsul lebih lanjut ke Psikiater"}
	Q[*P] = load_question(11, "Pusing", "Cara meredakan pusing akibat tugas?", list11, "1 Maret 2021 22:33", "Arie Farchan Fyrzatullah")
	Q[*P].jumlah_tanggapan = 3
	*P += 1

	list12 := []string{"dr.Rina : DPT, MMR, polio, dan hepatitis B."}
	Q[*P] = load_question(12, "Imun", "Apa saja vaksin yang direkomendasikan untuk anak-anak?", list12, "23 April 2018 15:12", "Rizki Nata")
	Q[*P].jumlah_tanggapan = 1
	*P += 1

	list13 := []string{"dr.Hadi : Debu, asap rokok, dan bulu hewan.", "Jundi Haq : Bagaimana cara mengatasi serangan asma?", "dr.Hadi : Gunakan inhaler sesuai anjuran dokter."}
	Q[*P] = load_question(13, "Asma", "Apa saja vaksin yang direkomendasikan untuk anak-anak?", list13, "11 Januari 2019 9:31", "Jundi Haq")
	Q[*P].jumlah_tanggapan = 3
	*P += 1
}
