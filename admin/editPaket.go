package admin

import (
	"PackageDelivery/datas"
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
	found := datas.CariPaket(noResi)
	if found.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
		return
	}
	paket.UbahPaket(found)
}