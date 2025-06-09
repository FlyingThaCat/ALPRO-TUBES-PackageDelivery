package paket

import (
	"PackageDelivery/utils"
	"fmt"
)

func CekPaket() {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("📦  CEK PAKET")
	fmt.Println("========================================")

	noResi := utils.GetString("Masukkan No Resi: ", "Silakan masukkan No Resi yang valid.")
	found := utils.FindPaketByNoResi(noResi)

	if found.NoResi == "" {
		fmt.Println("\n⚠️  Paket tidak ditemukan.")
	} else {
		fmt.Println()
		DetailPaket(*found)
	}

	utils.EnterToContinue()
}
