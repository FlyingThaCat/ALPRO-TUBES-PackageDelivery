package paket

import (
	"PackageDelivery/utils"
	"fmt"
)

func HapusPaket() {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("🗑️  HAPUS PAKET")
	fmt.Println("========================================")

	noResi := utils.GetString("Masukkan No Resi Paket yang ingin dihapus: ", "⚠️  Silakan masukkan No Resi yang valid.")

	ok := utils.DeletePaket(noResi)

	if !ok {
		fmt.Printf("\n⚠️  Paket dengan No Resi %s tidak ditemukan. Penghapusan dibatalkan.\n\n", noResi)
		utils.EnterToContinue()
		return
	}

	fmt.Printf("\n✅  Paket dengan No Resi %s telah berhasil dihapus.\n\n", noResi)
	utils.EnterToContinue()
}
