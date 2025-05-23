package admin

import (
	"PackageDelivery/utils"
	"fmt"
)

func HapusPaket() {
	var noResi string
	utils.ClearScreen()
	fmt.Print("=== Hapus Paket ===\n")
	fmt.Print("Masukkan No Resi Paket yang ingin dihapus: ")
	fmt.Scanln(&noResi)

	ok := utils.DeletePaket(noResi)

	if !ok {
		fmt.Println("Paket tidak dapat dihapus. No Resi tidak ditemukan.")
		return
	}

	fmt.Printf("Paket dengan No Resi %s telah dihapus.\n\n", noResi)

	utils.EnterToContinue()
}