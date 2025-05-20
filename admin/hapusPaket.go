package admin

import (
	"PackageDelivery/paket"
	"PackageDelivery/utils"
	"fmt"
)

func HapusPaket() {
	var noResi string
	fmt.Print("\n=== Hapus Paket ===\n")
	fmt.Print("Masukkan No Resi Paket yang ingin dihapus: ")
	fmt.Scanln(&noResi)

	ok := paket.HapusPaket(noResi)

	if !ok {
		fmt.Println("Paket tidak dapat dihapus. No Resi tidak ditemukan.")
		return
	}

	fmt.Printf("Paket dengan No Resi %s telah dihapus.\n\n", noResi)

	utils.EnterToContinue()
}