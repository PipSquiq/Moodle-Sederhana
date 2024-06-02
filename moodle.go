package main

import (
	"fmt"
)

// Struktur untuk menyimpan data komentar di forum
type Komentar struct {
	Nama     string
	Komentar string
}

// Variabel global untuk menyimpan forum
var forum []Komentar

// Struktur untuk menyimpan data soal dan jawaban
type Soal struct {
	Pertanyaan string
	Jawaban    string
}

// Struktur untuk menyimpan nilai Siswa
type Siswa struct {
	Nama       string
	NilaiTugas float64 // Menyimpan nilai tugas
	NilaiQuiz  float64 // Menyimpan nilai quiz
}

// Slice untuk menyimpan tugas dan quiz
var (
	tugasList []Soal
	quizList  []Soal
)

// Map untuk menyimpan nilai-nilai Siswa
var nilaiSiswaList = make(map[string]Siswa)

// Fungsi untuk menampilkan menu Guru
func menuGuru() {
	var pilihan int

	for {
		fmt.Println("\nMenu Guru:")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Ubah Tugas")
		fmt.Println("3. Hapus Tugas")
		fmt.Println("4. Tambah Quiz")
		fmt.Println("5. Ubah Quiz")
		fmt.Println("6. Hapus Quiz")
		fmt.Println("7. Lihat Nilai Siswa")
		fmt.Println("8. Lihat Forum")
		fmt.Println("9. Tambahkan Komentar")
		fmt.Println("10. Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahSoal(&tugasList)
		case 2:
			ubahSoal(&tugasList)
		case 3:
			hapusSoal(&tugasList)
		case 4:
			tambahSoal(&quizList)
		case 5:
			ubahSoal(&quizList)
		case 6:
			hapusSoal(&quizList)
		case 7:
			lihatNilaiSiswa()
		case 8:
			lihatForum()
		case 9:
			tambahKomentar("Guru")
		case 10:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk menampilkan menu Siswa
func menuSiswa(nama string) {
	var pilihan int

	for {
		fmt.Println("\nMenu Siswa:")
		fmt.Println("1. Jawab Tugas")
		fmt.Println("2. Jawab Quiz")
		fmt.Println("3. Lihat Forum")
		fmt.Println("4. Tambahkan Komentar")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			jawabSoal(&tugasList, nama, true)
		case 2:
			jawabSoal(&quizList, nama, false)
		case 3:
			lihatForum()
		case 4:
			tambahKomentar(nama)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk melihat forum
func lihatForum() {
	if len(forum) == 0 {
		fmt.Println("Forum masih kosong.")
		return
	}
	fmt.Println("\nForum:")
	for _, komentar := range forum {
		fmt.Printf("%s: %s\n", komentar.Nama, komentar.Komentar)
	}
}

// Fungsi untuk menambah komentar di forum
func tambahKomentar(nama string) {
	var komentar Komentar
	komentar.Nama = nama
	fmt.Print("Masukkan komentar: ")
	fmt.Scanln(&komentar.Komentar)
	forum = append(forum, komentar)
	fmt.Println("Komentar berhasil ditambahkan.")
}

// Fungsi untuk menambah soal
func tambahSoal(list *[]Soal) {
	var soal Soal
	fmt.Print("Masukkan pertanyaan: ")
	fmt.Scanln(&soal.Pertanyaan)
	*list = append(*list, soal)
	fmt.Println("Soal berhasil ditambahkan.")
	tambahJawaban(list)
}

// Fungsi untuk menambah jawaban
func tambahJawaban(list *[]Soal) {
	if len(*list) == 0 {
		fmt.Println("Belum ada soal yang tersedia.")
		return
	}
	lihatSoal(*list)
	var index int
	fmt.Print("Masukkan nomor soal yang ingin ditambahkan jawabannya: ")
	fmt.Scan(&index)
	if index > 0 && index <= len(*list) {
		fmt.Print("Masukkan jawaban: ")
		fmt.Scanln(&(*list)[index-1].Jawaban)
		fmt.Println("Jawaban berhasil ditambahkan.")
	} else {
		fmt.Println("Nomor soal tidak valid.")
	}
}

// Fungsi untuk mengubah soal
func ubahSoal(list *[]Soal) {
	var index int
	lihatSoal(*list)
	fmt.Print("Masukkan nomor soal yang ingin diubah: ")
	fmt.Scan(&index)
	if index > 0 && index <= len(*list) {
		fmt.Print("Masukkan pertanyaan baru: ")
		fmt.Scanln(&(*list)[index-1].Pertanyaan)
		fmt.Print("Masukkan jawaban baru: ")
		fmt.Scanln(&(*list)[index-1].Jawaban)
		fmt.Println("Soal berhasil diubah.")
	} else {
		fmt.Println("Nomor soal tidak valid.")
	}
}

// Fungsi untuk menghapus soal
func hapusSoal(list *[]Soal) {
	var index int
	lihatSoal(*list)
	fmt.Print("Masukkan nomor soal yang ingin dihapus: ")
	fmt.Scan(&index)
	if index > 0 && index <= len(*list) {
		*list = append((*list)[:index-1], (*list)[index:]...)
		fmt.Println("Soal berhasil dihapus.")
	} else {
		fmt.Println("Nomor soal tidak valid.")
	}
}

// Fungsi untuk melihat daftar soal
func lihatSoal(list []Soal) {
	for i, soal := range list {
		fmt.Printf("%d. %s\n", i+1, soal.Pertanyaan)
	}
}

// Fungsi untuk menjawab soal
func jawabSoal(list *[]Soal, nama string, isTugas bool) {
	if len(*list) == 0 {
		fmt.Println("Belum ada soal yang tersedia.")
		return
	}
	lihatSoal(*list)
	var index int
	var jawaban string
	var totalNilai float64
	totalSoal := float64(len(*list))

	fmt.Print("Masukkan nomor soal yang ingin dijawab: ")
	fmt.Scan(&index)
	if index > 0 && index <= len(*list) {
		fmt.Print("Masukkan jawaban: ")
		fmt.Scanln(&jawaban)
		if jawaban == (*list)[index-1].Jawaban {
			fmt.Println("Jawaban benar! Anda mendapatkan nilai untuk soal ini.")
			totalNilai = 100.0 / totalSoal
		} else {
			fmt.Printf("Jawaban salah. Jawaban yang benar adalah: %s\n", (*list)[index-1].Jawaban)
			totalNilai = 0.0
		}
	} else {
		fmt.Println("Nomor soal tidak valid.")
		return
	}

	if isTugas {
		tambahNilaiSiswa(nama, totalNilai, 0)
	} else {
		tambahNilaiSiswa(nama, 0, totalNilai)
	}
}

// Fungsi untuk menambahkan atau memperbarui nilai Siswa
func tambahNilaiSiswa(nama string, nilaiTugas, nilaiQuiz float64) {
	siswa, ok := nilaiSiswaList[nama]
	if !ok {
		siswa = Siswa{Nama: nama}
	}
	if nilaiTugas > 0 {
		siswa.NilaiTugas += nilaiTugas
		if siswa.NilaiTugas > 100 {
			siswa.NilaiTugas = 100
		}
	}
	if nilaiQuiz > 0 {
		siswa.NilaiQuiz += nilaiQuiz
		if siswa.NilaiQuiz > 100 {
			siswa.NilaiQuiz = 100
		}
	}
	nilaiSiswaList[nama] = siswa
}

// Fungsi untuk menampilkan nilai Siswa
func lihatNilaiSiswa() {
	if len(nilaiSiswaList) == 0 {
		fmt.Println("Belum ada nilai Siswa yang tersedia.")
		return
	}

	// Menampilkan nilai Siswa
	fmt.Println("\nNilai Siswa:")
	for _, siswa := range nilaiSiswaList {
		totalNilai := (siswa.NilaiTugas + siswa.NilaiQuiz) / 2
		fmt.Printf("%s - Nilai Tugas: %.2f\n", siswa.Nama, siswa.NilaiTugas)
		fmt.Printf("%s - Nilai Quiz: %.2f\n", siswa.Nama, siswa.NilaiQuiz)
		fmt.Printf("%s - Total Nilai: %.2f\n", siswa.Nama, totalNilai)
	}
}

func main() {
	for {
		var role, nama string

		// Meminta pengguna untuk memasukkan peran
		fmt.Println("Masuk Sebagai (Guru/Siswa):")
		fmt.Scanln(&role)

		if role == "Guru" {
			menuGuru()
		} else if role == "Siswa" {
			// Meminta pengguna untuk memasukkan nama jika perannya Siswa
			fmt.Println("Masukkan nama Anda:")
			fmt.Scanln(&nama)
			menuSiswa(nama)
		} else {
			fmt.Println("Peran tidak dikenal. Silakan jalankan kembali program dan masukkan peran yang valid.")
		}
	}
}
