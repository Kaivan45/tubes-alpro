package main

import (
	"fmt"
	"strings"
)

// Struct untuk pengguna
type User struct {
	Username string
	Password string
	Role     string
}

// Struct untuk pertanyaan
type Pertanyaan struct {
	ID        int
	Penanya   string
	Isi       string
	Tag       []string
	Tanggapan []string
}

var users []User
var forum []Pertanyaan
var currentUser *User

func registrasi() {
	var username, password, role string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)
	fmt.Print("Masukkan role (pasien/dokter): ")
	fmt.Scan(&role)

	// Validasi role
	if role != "pasien" && role != "dokter" {
		fmt.Println("Role tidak valid, hanya bisa pasien atau dokter.")
		return
	}

	// Simpan user baru
	users = append(users, User{Username: username, Password: password, Role: role})
	fmt.Println("Registrasi berhasil!")
}

func login() {
	var username, password string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for i := 0; i < len(users); i++ {
		if users[i].Username == username && users[i].Password == password {
			currentUser = &users[i]
			fmt.Printf("Login berhasil! Selamat datang, %s (%s)\n", users[i].Username, users[i].Role)
			return
		}
	}
	fmt.Println("Username atau password salah.")
}

func lihatForum() {
	if len(forum) == 0 {
		fmt.Println("Belum ada pertanyaan di forum.")
		return
	}

	fmt.Println("Daftar Pertanyaan:")
	for i := 0; i < len(forum); i++ {
		fmt.Printf("ID: %d | Penanya: %s\nPertanyaan: %s\nTag: %s\n",
			forum[i].ID, forum[i].Penanya, forum[i].Isi, strings.Join(forum[i].Tag, ", "))
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

	var isi string
	var tagInput string
	fmt.Print("Masukkan pertanyaan Anda: ")
	fmt.Scan(&isi)
	fmt.Print("Masukkan tag (pisahkan dengan koma): ")
	fmt.Scan(&tagInput)

	tags := strings.Split(tagInput, ",")
	pertanyaan := Pertanyaan{
		ID:        len(forum) + 1,
		Penanya:   currentUser.Username,
		Isi:       isi,
		Tag:       tags,
		Tanggapan: []string{},
	}
	forum = append(forum, pertanyaan)
	fmt.Println("Pertanyaan berhasil diposting!")
}

func main() {
	var pil int
	for {
		fmt.Println("\n===================================================")
		fmt.Println("= SELAMAT DATANG DI APLIKASI KONSULTASI KESEHATAN =")
		fmt.Println("===================================================")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Forum")
		if currentUser != nil && currentUser.Role == "pasien" {
			fmt.Println("4. Posting Pertanyaan")
		}
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scan(&pil)

		// Validasi input menu
		if currentUser == nil && pil == 4 {
			fmt.Println("Anda harus login sebagai pasien untuk mengakses menu ini.")
			continue
		}

		switch pil {
		case 1:
			registrasi()
		case 2:
			login()
		case 3:
			lihatForum()
		case 4:
			if currentUser != nil && currentUser.Role == "pasien" {
				postingPertanyaan()
			} else {
				fmt.Println("Menu ini hanya tersedia untuk pasien yang sudah login.")
			}
		case 5:

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
