package datas

import (
	"PackageDelivery/utils"
	"fmt"
)

func ListAllKurir() {
	utils.ClearScreen()
	fmt.Println("=== Daftar Kurir ===")
	for _, kurir := range UsersDB {
		if kurir.Role != "kurir" {
			continue
		}
		fmt.Printf("Username: %s\n", kurir.Username)
		fmt.Println("---------------------")
	}

	utils.EnterToContinue()
}

func HapusKurir(username string) bool {
	// Hapus Kurir dari database
	for i, p := range UsersDB {
		if p.Username == username {
			UsersDB = append(UsersDB[:i], UsersDB[i+1:]...)
			return true
		}
	}
	return false
}