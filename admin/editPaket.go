package admin

import (
	"PackageDelivery/paket"
	"PackageDelivery/utils"
	"fmt"
)

func EditPaket() {
	utils.ClearScreen()
	fmt.Println("=== Edit Paket ===")
	
	noResi := utils.GetString("Masukkan No Resi Paket yang ingin diedit: ", "Silakan masukkan No Resi yang valid.")
	
	found := utils.FindPaketByNoResi(noResi)
	if found.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
		return
	}
	
	paket.UbahPaket(*found)
}