package paket

import (
	"PackageDelivery/utils"
	"fmt"
)

func HapusPaket() {
	utils.ClearScreen()
	fmt.Print("=== Hapus Paket ===\n")
	
	noResi := utils.GetString("Masukkan No Resi Paket yang ingin dihapus: ", "Silakan masukkan No Resi yang valid.")

	ok := utils.DeletePaket(noResi)

	if !ok {
		fmt.Println("Paket tidak dapat dihapus. No Resi tidak ditemukan.")
		return
	}

	fmt.Printf("Paket dengan No Resi %s telah dihapus.\n\n", noResi)

	utils.EnterToContinue()
}