package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"PackageDelivery/admin"
	"PackageDelivery/kurir"
	"PackageDelivery/types"
	"PackageDelivery/users"

	"PackageDelivery/datas"
)


func register() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Register User ===")
	fmt.Print("Masukkan username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if _, exists := datas.UsersDB[username]; exists {
		fmt.Println("Username sudah ada, silakan coba lagi.\n")
		return
	}

	fmt.Print("Masukkan password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Masukkan role (user/kurir/admin): ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	if role != "user" && role != "kurir" && role != "admin" {
		fmt.Println("Role tidak valid. Harus 'user', 'kurir', atau 'admin'.\n")
		return
	}

	datas.UsersDB[username] = types.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	fmt.Printf("User '%s' dengan role '%s' berhasil didaftarkan.\n\n", username, role)
}

func login() (*types.User, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Login ===")
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user, ok := datas.UsersDB[username]
	if !ok || user.Password != password {
		fmt.Println("Username atau password salah.\n")
		return nil, false
	}

	fmt.Printf("Login berhasil, selamat datang %s! (role: %s)\n\n", user.Username, user.Role)
	return &user, true
}

func main() {
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
						fmt.Println("Logout berhasil.\n")
						break
					} else {
						fmt.Printf("Menu pilihan '%s' sedang dalam pengembangan.\n\n", opt)
					}
				}
			}
		case "3":
			fmt.Println("Terima kasih, sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.\n")
		}
	}
}
