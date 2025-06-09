package main

import (
	"PackageDelivery/paket"
	"PackageDelivery/utils"
	"fmt"

	"PackageDelivery/admin"
	"PackageDelivery/kurir"
	"PackageDelivery/types"
	"PackageDelivery/users"

	"PackageDelivery/datas"
)

func register() {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("🆕  REGISTER USER")
	fmt.Println("========================================")

	username := utils.GetString("Masukkan username: ", "Username tidak boleh kosong.")

	if user := utils.FindUserByUsername(username); user.Username != "" {
		fmt.Println("\n❌ Username sudah ada, silakan coba lagi.\n")
		return
	}

	password := utils.GetString("Masukkan password: ", "Password tidak boleh kosong.")

	utils.AddUser(types.User{
		Username: username,
		Password: password,
		Role:     "user",
	})

	fmt.Printf("\n✅ User '%s' berhasil didaftarkan.\n\n", username)
}

func login() (*types.User, bool) {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("🔐  LOGIN")
	fmt.Println("========================================")

	username := utils.GetString("Masukkan username: ", "Username tidak boleh kosong.")
	password := utils.GetString("Masukkan password: ", "Password tidak boleh kosong.")

	user := utils.FindUserByUsername(username)
	if user.Username == "" {
		utils.ShowDelayedMessage("\n❌ Username tidak ditemukan.\n", 2, true)
		return nil, false
	} else if user.Password != password {
		utils.ShowDelayedMessage("\n❌ Password salah.\n", 2, true)
		return nil, false
	}

	if user.Role != "user" {
		fmt.Printf("\n✅ Login berhasil, selamat datang %s! (role: %s)\n\n", user.Username, user.Role)
	} else {
		fmt.Printf("\n✅ Login berhasil, selamat datang %s!\n\n", user.Username)
	}

	utils.SetLoggedInUsername(user.Username)
	return &user, true
}

func main() {
	utils.ClearScreen()
	datas.Init() // Inisialisasi data awal

	for {
		utils.ClearScreen()
		fmt.Println("========================================")
		fmt.Println("🔐  SELAMAT DATANG DI APLIKASI")
		fmt.Println("========================================")
		fmt.Println("Silakan pilih menu:")
		fmt.Println("1️⃣  Register")
		fmt.Println("2️⃣  Login")
		fmt.Println("0️⃣  Keluar")
		fmt.Println("========================================")
		choice := utils.GetInt("Pilih menu: ", "")

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
								kurir.CheckMyPaketSorted()
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
