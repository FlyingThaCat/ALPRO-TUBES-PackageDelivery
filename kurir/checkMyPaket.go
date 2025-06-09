package kurir

import (
	"PackageDelivery/utils"
	"fmt"
)

func CheckMyPaket() {
	utils.ClearScreen()

	username := utils.GetLoggedInUsername()

	pakets := utils.FindPaketByUsername(username)

	if len(pakets) == 0 {
		fmt.Println("========================================")
		fmt.Println("📦  DAFTAR PAKET ANDA")
		fmt.Println("========================================")
		fmt.Print("\nTidak ada paket yang sedang Anda tangani.\n")
		utils.EnterToContinue()
		return
	}

	fmt.Println("========================================")
	fmt.Println("📦  DAFTAR PAKET ANDA")
	fmt.Println("========================================")

	for i, paket := range pakets {
		fmt.Printf("%d. No Resi: %s | Tipe: %s | Dari %s ke %s | Status Terakhir: %s\n",
			i+1, paket.NoResi, paket.Tipe, paket.SenderCity, paket.ReceiverCity,
			paket.Status[len(paket.Status)-1],
		)
	}

	fmt.Println("========================================")
	utils.EnterToContinue()
}
