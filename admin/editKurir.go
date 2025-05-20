package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/kurir"
	"PackageDelivery/utils"
	"fmt"
)

func EditKurir() {
	utils.ClearScreen()
	fmt.Println("=== Edit Kurir ===")
	fmt.Print("Masukkan Username Kurir yang ingin diedit: ")
	var username string
	fmt.Scanln(&username)
	found := datas.FindUserByUsername(username)
	if found.Username == "" {
		fmt.Println("Kurir tidak ditemukan.")
		return
	}
	kurir.UbahKurir(found)
}