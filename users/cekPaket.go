package users

import (
	"PackageDelivery/paket"
	"fmt"
)

func CekPaket() {
	var noResi string
	fmt.Println("=== Cek Paket ===")
	fmt.Print("Masukkan No Resi: ")
	fmt.Scanln(&noResi)
	found := paket.CariPaket(noResi)
	if found.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
	} else {
		paket.DetailPaket(found)
	}
}