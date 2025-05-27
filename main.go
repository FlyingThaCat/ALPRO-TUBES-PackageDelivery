package main

import (
	"fmt"
	"PackageDelivery/paket"
	"PackageDelivery/utils"

	"PackageDelivery/admin"
	"PackageDelivery/kurir"
	"PackageDelivery/types"
	"PackageDelivery/users"

	"PackageDelivery/datas"
)

func register() {
	utils.ClearScreen()
	fmt.Print("=== Register User ===\n")
	username := utils.GetString("Masukkan username: ", "username tidak boleh kosong.")

	if user := utils.FindUserByUsername(username); user.Username != "" {
		fmt.Print("Username sudah ada, silakan coba lagi.\n\n")
		return
	}

	password := utils.GetString("Masukkan password: ", "password tidak boleh kosong.")

	utils.AddUser(types.User{
		Username: username,
		Password: password,
		Role:     "user", 
	})

	fmt.Printf("User '%s' berhasil didaftarkan.\n\n", username)
}

func login() (*types.User, bool) {
	utils.ClearScreen()
	fmt.Printf("=== Login ===\n")

	username := utils.GetString("Masukkan username: ", "username tidak boleh kosong.")
	password := utils.GetString("Masukkan password: ", "password tidak boleh kosong.")

	user := utils.FindUserByUsername(username)
	if user.Username == "" {
		utils.ShowDelayedMessage("Username tidak ditemukan.", 2, true)
		return nil, false
	} else if user.Password != password {
		utils.ShowDelayedMessage("Password salah.", 2, true)
		return nil, false
	}

	// BUGS
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

	for {
		utils.ClearScreen()
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")
		choice := utils.GetInt("Masukkan pilihan (1/2/0): ", "")

		switch choice {
		case 1:
			register()
		case 2:
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
					
					option := utils.GetInt("Pilih menu: ", "Silakan masukkan menu yang valid (0 untuk logout).")
					if option == 0 {
						fmt.Print("Logout berhasil.\n\n")
						break
					} else {
						switch user.Role {
						case "user":
							if option == 1 {
								paket.TambahPaket()
							} else if option == 2 {
								paket.CekPaket()
							} else {
								users.MenuUser()
							}
						case "admin":
							if option == 1 {
								paket.TambahPaket()
							} else if option == 2 {
								paket.LihatPaket(true, false)
							} else if option == 3 {
								paket.EditPaket()
							} else if option == 4 {
								paket.HapusPaket()
							} else if option == 5 {
								admin.TambahKurir()
							} else if option == 6 {
								admin.LihatKurir(true)
							} else if option == 7 {
								admin.EditKurir()
							} else if option == 8 {
								admin.HapusKurir()
							} else if option == 9 {
								paket.AssignPaket()
							} else {
								admin.MenuAdmin()
							}
						case "kurir":
							if option == 1 {
								kurir.CheckMyPaket()
							} else if option == 2 {
								kurir.UpdateStatus()
							} else {
								kurir.MenuKurir()
							}
						}
					}
				}
			}
		case 0:
			fmt.Println("Terima kasih, sampai jumpa!")
			return
		default:
			utils.ShowDelayedMessage("Pilihan tidak valid.\n\nSilakan masukkan 1, 2, atau 0 untuk keluar.", 2, true)
		}
	}
}
