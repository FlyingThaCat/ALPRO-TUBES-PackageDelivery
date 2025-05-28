package users

import (
	"PackageDelivery/utils"
	"fmt"
	"strings"
)

func MenuUser() {
	utils.ClearScreen()

	username := utils.GetLoggedInUsername()

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("📦  SELAMAT DATANG, %s\n", strings.ToUpper(username))
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("Silakan pilih menu:")
	fmt.Println("1️⃣  Tambah Paket")
	fmt.Println("2️⃣  Cek Paket")
	fmt.Println("0️⃣  Logout")
	fmt.Println(strings.Repeat("=", 40))
}
