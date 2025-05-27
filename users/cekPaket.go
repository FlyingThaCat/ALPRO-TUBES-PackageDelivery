package users

import (
	"PackageDelivery/paket"
	"PackageDelivery/utils"
	"fmt"
)

func CekPaket() {
	utils.ClearScreen()
	fmt.Println("=== Cek Paket ===")
	
	noResi := utils.GetInput("Masukkan No Resi: ")
	found := utils.FindPaketByNoResi(noResi)
	if found.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
	} else {
		paket.DetailPaket(*found)
	}
}