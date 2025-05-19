package users

import (
	"PackageDelivery/utils"
	"fmt"
)

func MenuUser() {
	utils.ClearScreen()

	fmt.Println("=== Menu User ===")
	fmt.Println("1. Tambah Paket")
	fmt.Println("2. Cek Paket")
	fmt.Println("0. Logout")
}