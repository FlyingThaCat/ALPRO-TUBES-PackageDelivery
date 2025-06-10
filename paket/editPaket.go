package paket

import (
	"PackageDelivery/utils"
	"fmt"
)

func EditPaket() {
	utils.ClearScreen()

	fmt.Println("========================================")
	fmt.Println("✏️  EDIT PAKET")
	fmt.Println("========================================")

	noResi := utils.GetString("Masukkan No Resi paket yang ingin diedit: ", "Silakan masukkan No Resi yang valid.")

	found := utils.FindPaketByNoResi(noResi)
	if found == nil {
		fmt.Println("⚠️  Paket tidak ditemukan.")
		utils.EnterToContinue()
		return
	}

	UbahPaket(*found)
}
