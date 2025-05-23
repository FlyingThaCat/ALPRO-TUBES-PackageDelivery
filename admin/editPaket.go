package admin

import (
	"PackageDelivery/paket"
	"PackageDelivery/utils"
	"fmt"
)

func EditPaket() {
	utils.ClearScreen()
	fmt.Println("=== Edit Paket ===")
	fmt.Print("Masukkan No Resi Paket yang ingin diedit: ")
	var noResi string
	fmt.Scanln(&noResi)
	found := utils.FindPaketByNoResi(noResi)
	if found.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
		return
	}
	paket.UbahPaket(*found)
}