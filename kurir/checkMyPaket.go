package kurir

import (
	"PackageDelivery/utils"
	"fmt"
)

func CheckMyPaket() {
	utils.ClearScreen()

	username := utils.GetLoggedInUsername()

	pakets := (utils.FindPaketByUsername(username))

	if len(pakets) == 0 {

		fmt.Print("\nTidak ada paket yang sedang Anda tangani.\n")
		utils.EnterToContinue()
		return
	}

	// utils.EnterToContinue()
	fmt.Println("=== Daftar Paket Anda ===")
	for i, paket := range pakets {
		fmt.Printf("%d. No Resi: %s | Tipe: %s | Dari %s ke %s | Status Terakhir: %s\n",
			i+1, paket.NoResi, paket.Tipe, paket.SenderCity, paket.ReceiverCity,
			paket.Status[len(paket.Status)-1],
		)
	}
	utils.EnterToContinue()

}
