package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Deklarasi Array tetap dengan kapasitasi 100
const maxUsers = 100
const maxForum = 100

// Deklarasi Struct untuk menyimpan informasi user
type User struct {
	Username string
	Password string
	Role     string
}

// Deklarasi Struct untuk menyimpan pertanyaan
type Pertanyaan struct {
	ID        int
	Penanya   string
	Isi       string
	Tag       []string
	Tanggapan []string
}

var users [maxUsers]User
var userCount int
var forum [maxForum]Pertanyaan
var forumCount int

var currentUser *User

func registrasi() {
	var username, password, role string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)
	fmt.Print("Masukkan role (pasien/dokter): ")
	fmt.Scan(&role)

	// validasi role pasien dan dokter
	if role != "pasien" && role != "dokter" {
		fmt.Println("Role tidak valid, hanya bisa pasien atau dokter.")
		return
	}

	// Meriksa ruang di Array
	if userCount >= maxUsers {
		fmt.Println("Registrasi gagal. Kapasitas pengguna penuh.")
		return
	}

	// Menyimpan user baru ke Array
	users[userCount] = User{Username: username, Password: password, Role: role}
	userCount++
	fmt.Println("Registrasi berhasil!")
}

func login() {
	var username, password string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	// Mencari user dengan username dan password yang sesuai
	for i := 0; i < userCount; i++ {
		if users[i].Username == username && users[i].Password == password {
			currentUser = &users[i]
			fmt.Printf("Login berhasil! Selamat datang, %s (%s)\n", users[i].Username, users[i].Role)
			return
		}
	}
	fmt.Println("Username atau password salah.")
}

func logout() {
	if currentUser != nil {
		fmt.Printf("Logout berhasil! Sampai jumpa, %s.\n", currentUser.Username)
		currentUser = nil
	} else {
		fmt.Println("Anda belum login.")
	}
}

func lihatForum() {
	if forumCount == 0 {
		fmt.Println("Belum ada pertanyaan di forum.")
		return
	}

	// Memilih algoritma pengurutan (Insertion Sort atau Selection Sort)
	var pil int
	fmt.Println("Pilih metode pengurutan berdasarkan tag:")
	fmt.Println("1. Ascending (Insertion Sort)")
	fmt.Println("2. Descending (Selection Sort)")
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scan(&pil)

	// Urutkan menggunakan metode yang dipilih
	if pil == 1 {
		insertionSortTag()
	} else if pil == 2 {
		selectionSortTagDescending()
	} else {
		fmt.Println("Pilihan tidak valid. Menampilkan forum tanpa pengurutan.")
	}

	// Tampilkan daftar pertanyaan yang sudah diurutkan
	fmt.Println("Daftar Pertanyaan:")
	for i := 0; i < forumCount; i++ {
		fmt.Printf("Urutan: %d | Penanya: %s\nPertanyaan: %s\nTag: %s\n",
			i+1, forum[i].Penanya, forum[i].Isi, strings.Join(forum[i].Tag, ", "))
		if len(forum[i].Tanggapan) > 0 {
			fmt.Println("Tanggapan:")
			for j := 0; j < len(forum[i].Tanggapan); j++ {
				fmt.Printf("- %s\n", forum[i].Tanggapan[j])
			}
		}
		fmt.Println("-----------------------------")
	}
}

func postingPertanyaan() {
	if currentUser == nil || currentUser.Role != "pasien" {
		fmt.Println("Hanya pasien yang dapat memposting pertanyaan.")
		return
	}

	if forumCount >= maxForum {
		fmt.Println("Forum penuh. Tidak bisa memposting pertanyaan baru.")
		return
	}

	tanyaReader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan pertanyaan Anda (akhiri dengan enter): ")
	isiPertanyaan, errPertanyaan := tanyaReader.ReadString('\n')
	if errPertanyaan != nil {
		fmt.Println("Terjadi kesalahan input pertanyaan.")
		return
	}
	isiPertanyaan = strings.TrimSpace(isiPertanyaan)
	if isiPertanyaan == "" {
		fmt.Println("Pertanyaan tidak boleh kosong.")
		return
	}
	isiPertanyaan = strings.TrimSpace(isiPertanyaan)

	fmt.Print("Masukkan tag: ")
	tagInput, errTags := tanyaReader.ReadString('\n')
	if errTags != nil {
		fmt.Println("Terjadi kesalahan input tag.")
		return
	}
	tagInput = strings.TrimSpace(tagInput)

	tags := strings.Split(tagInput, ",")

	maxTanggapan := 100
	tanggapan := make([]string, 0, maxTanggapan)

	pertanyaan := Pertanyaan{
		ID:        forumCount + 1,
		Penanya:   currentUser.Username,
		Isi:       isiPertanyaan,
		Tag:       tags,
		Tanggapan: tanggapan,
	}

	forum[forumCount] = pertanyaan
	forumCount++
	fmt.Println("Pertanyaan berhasil diposting!")
}

func beriTanggapan() {
	if currentUser == nil {
		fmt.Println("Anda harus login untuk memberikan tanggapan.")
		return
	}

	if forumCount == 0 {
		fmt.Println("Tidak ada pertanyaan di forum untuk ditanggapi.")
		return
	}

	fmt.Println("Daftar Pertanyaan:")
	for i := 0; i < forumCount; i++ {
		fmt.Printf("Urutan: %d | Penanya: %s\nPertanyaan: %s\nTag: %s\n",
			i+1, forum[i].Penanya, forum[i].Isi, strings.Join(forum[i].Tag, ", "))
		if len(forum[i].Tanggapan) > 0 {
			fmt.Println("Tanggapan:")
			for j := 0; j < len(forum[i].Tanggapan); j++ {
				fmt.Printf("- %s\n", forum[i].Tanggapan[j])
			}
		}
		fmt.Println("-----------------------------")
	}

	var urutan int
	fmt.Print("Masukkan urutan pertanyaan yang ingin Anda tanggapi: ")
	fmt.Scan(&urutan)

	if urutan < 1 || urutan > forumCount {
		fmt.Println("Urutan pertanyaan tidak valid.")
		return
	}

	tanggapanReader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan tanggapan Anda: ")
	tanggapanUser, errInput := tanggapanReader.ReadString('\n')
	if errInput != nil {
		fmt.Println("Kesalahan pembacaan input tanggapan.")
		return
	}
	tanggapanUser = strings.TrimSpace(tanggapanUser)

	tanggapanFormatted := fmt.Sprintf("%s (%s): %s", currentUser.Username, currentUser.Role, tanggapanUser)

	// Tentukan kapasitas maksimum untuk tanggapan
	const maxTanggapan = 50

	// Jika ruang masih tersedia di array Tanggapan, tambahkan tanggapan
	if len(forum[urutan-1].Tanggapan) >= maxTanggapan {
		fmt.Println("Tanggapan tidak dapat ditambahkan, ruang sudah penuh.")
		return
	}

	// Array statis untuk menyimpan tanggapan, dengan kapasitas tetap
	var tanggapanArray [maxTanggapan]string

	// Copy tanggapan yang ada ke array statis
	copy(tanggapanArray[:], forum[urutan-1].Tanggapan)

	// Tambahkan tanggapan pada posisi kosong di array tanpa menggunakan break
	added := false
	for i := 0; i < maxTanggapan; i++ {
		if tanggapanArray[i] == "" { // Menemukan posisi kosong
			tanggapanArray[i] = tanggapanFormatted
			added = true
		}
	}

	// Jika tanggapan berhasil ditambahkan, salin kembali ke forum
	if added {
		forum[urutan-1].Tanggapan = tanggapanArray[:len(forum[urutan-1].Tanggapan)+1]
		fmt.Println("Tanggapan berhasil ditambahkan!")
	} else {
		fmt.Println("Gagal menambahkan tanggapan.")
	}
}

func cariPertanyaan() {
	var pil int
	fmt.Println("Pilih metode pencarian berdasarkan tag:")
	fmt.Println("1. Squential")
	fmt.Println("2. Binary")
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scan(&pil)

	// Urutkan menggunakan metode yang dipilih
	if pil == 1 {
		var tag string
		fmt.Print("Masukkan tag yang ingin dicari: ")
		fmt.Scan(&tag)
		cariPertanyaanSequential(tag)
	} else if pil == 2 {
		var tag string
		fmt.Print("Masukkan tag yang ingin dicari: ")
		fmt.Scan(&tag)
		cariPertanyaanBinary(tag)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func cariPertanyaanSequential(tag string) {
	fmt.Println("Hasil Pencarian Sequential:")
	found := false
	index := 0

	for index < len(forum) {
		pertanyaan := forum[index]
		tagIndex := 0
		isMatched := false

		for tagIndex < len(pertanyaan.Tag) && !isMatched {
			// Mengubah logika pencocokan untuk mendukung pencarian parsial
			isMatched = strings.Contains(strings.ToLower(pertanyaan.Tag[tagIndex]), strings.ToLower(tag))
			tagIndex++
		}

		if isMatched {
			fmt.Printf("ID: %d | Penanya: %s\nPertanyaan: %s\nTag: %s\n",
				pertanyaan.ID, pertanyaan.Penanya, pertanyaan.Isi, strings.Join(pertanyaan.Tag, ", "))
			found = true
		}
		index++
	}

	if !found {
		fmt.Println("Tidak ada pertanyaan dengan tag tersebut.")
	}
}

func cariPertanyaanBinary(tag string) {
	// Urutkan data
	insertionSortTag()

	low := 0
	high := forumCount - 1
	found := false

	fmt.Println("Hasil Pencarian Binary:")

	// Ubah tag menjadi huruf kecil untuk pencarian yang case-insensitive
	searchTag := strings.ToLower(tag)

	// Proses pencarian binary
	for low <= high {
		mid := (low + high) / 2
		// Ambil tag yang ada dalam pertanyaan
		forumTag := strings.ToLower(strings.Join(forum[mid].Tag, " ")) // Gabungkan semua tag menjadi satu string

		// Cek apakah tag mengandung searchTag
		if strings.Contains(forumTag, searchTag) {
			// Jika ditemukan, tampilkan data
			fmt.Printf("ID: %d | Penanya: %s\nPertanyaan: %s\nTag: %s\n",
				forum[mid].ID, forum[mid].Penanya, forum[mid].Isi, strings.Join(forum[mid].Tag, ", "))
			found = true
			break
		} else if forumTag < searchTag {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ada pertanyaan dengan tag tersebut.")
	}
}

func insertionSortTag() {
	for i := 1; i < forumCount; i++ {
		pertanyaan := forum[i]
		j := i - 1
		// Urutkan berdasarkan tag pertama (Anda bisa menyesuaikan logika untuk lebih dari satu tag jika diperlukan)
		for j >= 0 && strings.Compare(forum[j].Tag[0], pertanyaan.Tag[0]) > 0 {
			forum[j+1] = forum[j]
			j--
		}
		forum[j+1] = pertanyaan
	}
}

func selectionSortTagDescending() {
	for i := 0; i < forumCount-1; i++ {
		maxIdx := i
		// Cari elemen dengan tag terbesar di sisa array
		for j := i + 1; j < forumCount; j++ {
			if strings.Compare(forum[j].Tag[0], forum[maxIdx].Tag[0]) > 0 {
				maxIdx = j
			}
		}
		// Tukar elemen
		if maxIdx != i {
			forum[i], forum[maxIdx] = forum[maxIdx], forum[i]
		}
	}
}

func initDummyData() {
	// Data dummy pengguna
	users[userCount] = User{Username: "pasien1", Password: "password1", Role: "pasien"}
	userCount++
	users[userCount] = User{Username: "dokter1", Password: "password2", Role: "dokter"}
	userCount++

	// Data dummy pertanyaan
	forum[forumCount] = Pertanyaan{
		ID:      forumCount + 1,
		Penanya: "pasien1",
		Isi:     "Bagaimana cara mengobati sakit kepala tanpa obat?",
		Tag:     []string{"sakit", "kepala"},
	}
	forumCount++

	forum[forumCount] = Pertanyaan{
		ID:      forumCount + 1,
		Penanya: "pasien1",
		Isi:     "Apakah olahraga baik untuk penderita diabetes?",
		Tag:     []string{"diabetes", "olahraga"},
	}
	forumCount++
}

func main() {
	var pil int
	initDummyData()
	for {
		fmt.Println("\n===================================================")
		fmt.Println("= SELAMAT DATANG DI APLIKASI KONSULTASI KESEHATAN =")
		fmt.Println("===================================================")
		if currentUser == nil {
			fmt.Println("1. Registrasi")
			fmt.Println("2. Login")
			fmt.Println("3. Lihat Forum")
			fmt.Println("0. Keluar")
			fmt.Print("Masukkan pilihan Anda: ")
		} else {
			fmt.Printf("Selamat datang, %s (%s)\n", currentUser.Username, currentUser.Role)
			fmt.Println("1. Lihat Forum")
			fmt.Println("2. Logout")
			if currentUser.Role == "pasien" {
				fmt.Println("3. Posting Pertanyaan")
			}
			fmt.Println("4. Beri Tanggapan pada Pertanyaan")
			fmt.Println("5. Cari Teks Pertanyaan")
			fmt.Print("Masukkan pilihan Anda: ")
		}
		fmt.Scan(&pil)

		if currentUser == nil && pil == 4 {
			fmt.Println("Anda harus login sebagai pasien untuk mengakses menu ini.")
			continue
		}

		switch {
		case currentUser == nil && pil == 1:
			registrasi()
		case currentUser == nil && pil == 2:
			login()
		case currentUser == nil && pil == 3:
			lihatForum()
		case currentUser != nil && pil == 1:
			lihatForum()
		case currentUser != nil && pil == 2:
			logout()
		case currentUser != nil && currentUser.Role == "pasien" && pil == 3:
			postingPertanyaan()
		case currentUser != nil && pil == 4:
			beriTanggapan()
		case currentUser != nil && pil == 5:
			cariPertanyaan()
		case currentUser == nil && pil == 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
