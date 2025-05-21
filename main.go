package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"PackageDelivery/paket"
	"PackageDelivery/utils"

	"PackageDelivery/admin"
	"PackageDelivery/kurir"
	"PackageDelivery/types"
	"PackageDelivery/users"

	"PackageDelivery/datas"
)

var DEBUG_ADMIN = false
var DEBUG_KURIR = false
var DEBUG_USER = false

func register() {
	utils.ClearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("=== Register User ===\n")
	fmt.Print("Masukkan username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if user := datas.FindUserByUsername(username); user.Username != "" {
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
	utils.ClearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("=== Login ===\n")
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := datas.FindUserByUsername(username)
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

	utils.SetLoggedInUsername(user.Username)

	return &user, true
}

func main() {
	utils.ClearScreen()
	datas.Init() // Inisialisasi data awal
	reader := bufio.NewReader(os.Stdin)

	for {

		// utils.ClearScreen()
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan pilihan (1/2/0): ")
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

					if opt == "0" {
						fmt.Print("Logout berhasil.\n\n")
						break
					} else {

						switch user.Role {
						case "user":
							if opt == "1" {
								paket.TambahPaket()
							} else if opt == "2" {
								users.CekPaket()
							} else {
								fmt.Print("Pilihan tidak valid.\n\n")
							}
						case "admin":
							if opt == "1" {
								paket.TambahPaket()
							} else if opt == "2" {
								paket.LihatPaket(true)
							} else if opt == "3" {
								admin.EditPaket()
							} else if opt == "4" {
								admin.HapusPaket()
							} else if opt == "5" {
								admin.TambahKurir()
							} else if opt == "6" {
								admin.LihatKurir(true)
							} else if opt == "7" {
								admin.EditKurir()
							} else if opt == "8" {
								admin.HapusKurir()
							} else if opt == "9" {
								admin.AssignPaket()
							} else {
								fmt.Print("Pilihan tidak valid.\n\n")
							}
						case "kurir":
							if opt == "1" {
								kurir.CheckMyPaket()
							} else if opt == "2" {
								kurir.UpdateStatus()
							} else {
								fmt.Print("Pilihan tidak valid.\n\n")
							}
						}
					}
				}
			}
		case "0":
			fmt.Println("Terima kasih, sampai jumpa!")
			return
		default:
			fmt.Print("Pilihan tidak valid.\n\n")
		}
	}
}
