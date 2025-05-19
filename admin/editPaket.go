package admin

import (
	"PackageDelivery/paket"
	"fmt"
)

func EditPaket() {
	fmt.Println("=== Edit Paket ===")
	fmt.Print("Masukkan No Resi Paket yang ingin diedit: ")
	var noResi string
	fmt.Scanln(&noResi)
	found := paket.CariPaket(noResi)
	if found.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
		return
	}
	paket.UbahPaket(found)
}