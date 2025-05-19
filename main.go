package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"PackageDelivery/utils"

	"PackageDelivery/admin"
	"PackageDelivery/kurir"
	"PackageDelivery/types"
	"PackageDelivery/users"

	"PackageDelivery/datas"
)

func register() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n=== Register User ===\n")
	fmt.Print("Masukkan username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if user := utils.FindUserByUsername(username); user.Username != "" {
		fmt.Print("Username sudah ada, silakan coba lagi.\n\n")
		return
	}

	fmt.Print("Masukkan password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	datas.UsersDB = append(datas.UsersDB, types.User{
		Username: username,
		Password: password,
		Role:     "user",
	})

	fmt.Printf("User '%s' berhasil didaftarkan.\n\n", username)
}

func login() (*types.User, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n=== Login ===\n")
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := utils.FindUserByUsername(username)
	if user.Username == "" {
		fmt.Print("Username tidak ditemukan.\n\n")
		return nil, false
	} else if user.Password != password {
		fmt.Print("Password salah.\n\n")
		return nil, false
	}

	if user.Role != "user" { 
		fmt.Printf("Login berhasil, selamat datang %s! (role: %s)\n\n", user.Username, user.Role)
	} else {
		fmt.Printf("Login berhasil, selamat datang %s!\n\n", user.Username)
	}

	return &user, true
}

func main() {
	datas.Init() // Inisialisasi data awal
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan pilihan (1/2/3): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			register()
		case "2":
			user, loggedIn := login()
			if loggedIn {
				for {
					switch user.Role {
					case "user":
						users.MenuUser()
					case "kurir":
						kurir.MenuKurir()
					case "admin":
						admin.MenuAdmin()
					}

					fmt.Print("Pilih menu: ")
					opt, _ := reader.ReadString('\n')
					opt = strings.TrimSpace(opt)

					if opt == "3" {
						fmt.Print("Logout berhasil.\n\n")
						break
					} else {
						switch user.Role {
						case "user":
							if opt == "1" {
								users.TambahPaket()
							} else if opt == "2" {
								fmt.Print("Cek Paket belum tersedia.\n\n")
							} else {
								fmt.Print("Pilihan tidak valid.\n\n")
							}
						}
					}
				}
			}
		case "3":
			fmt.Println("Terima kasih, sampai jumpa!")
			return
		default:
			fmt.Print("Pilihan tidak valid.\n\n")
		}
	}
}
