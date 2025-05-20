package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/utils"
	"fmt"
)

func HapusKurir() {
	var username string
	utils.ClearScreen()
	fmt.Print("=== Hapus Kurir ===\n")
	fmt.Print("Masukkan username kurir yang ingin dihapus: ")
	fmt.Scanln(&username)

	ok := datas.HapusKurir(username)

	if !ok {
		fmt.Println("Kurir tidak dapat dihapus. Username tidak ditemukan.")
		return
	}

	fmt.Printf("Kurir dengan username %s telah dihapus.\n\n", username)

	utils.EnterToContinue()
}