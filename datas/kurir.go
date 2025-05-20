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