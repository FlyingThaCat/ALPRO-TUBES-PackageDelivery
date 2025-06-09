package admin

import (
	"PackageDelivery/kurir"
	"PackageDelivery/utils"
	"fmt"
)

func EditKurir() {
	utils.ClearScreen()
	fmt.Println("=== Edit Kurir ===")
	
	username := utils.GetString("Masukkan Username Kurir yang ingin diedit: ", "Silakan masukkan username kurir yang valid.")
	
	found := utils.FindUserByUsername(username)
	if found.Username == "" {
		fmt.Println("Kurir tidak ditemukan.")
		return
	}
	
	kurir.UbahKurir(found)
}
