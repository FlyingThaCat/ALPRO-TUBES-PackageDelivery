package users

import (
	"PackageDelivery/datas"
	"PackageDelivery/paket"
	"PackageDelivery/utils"
	"fmt"
)

func CekPaket() {
	utils.ClearScreen()
	var noResi string
	fmt.Println("=== Cek Paket ===")
	fmt.Print("Masukkan No Resi: ")
	fmt.Scanln(&noResi)
	found := datas.CariPaket(noResi)
	if found.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
	} else {
		paket.DetailPaket(found)
	}
}