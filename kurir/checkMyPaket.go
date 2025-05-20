package kurir

import (
	"PackageDelivery/datas"
	"PackageDelivery/utils"
	"fmt"
)

func CheckMyPaket() {
	utils.ClearScreen()

	username := utils.GetLoggedInUsername()

	fmt.Printf("Selamat datang kembali, kurir %s!\n", username)

	user := datas.FindUserByUsername(username)
	if user.Username == "" {

		fmt.Print("\nUser tidak ditemukan.\n")
		utils.EnterToContinue()
		return
	}

	if user.Role != "kurir" {

		fmt.Print("\nHanya kurir yang dapat mengakses fitur ini.\n")
		utils.EnterToContinue()
		return
	}

	if len(user.Pakets) == 0 {

		fmt.Print("\nTidak ada paket yang sedang Anda tangani.\n")
		utils.EnterToContinue()
		return
	}

	fmt.Println("=== Daftar Paket Anda ===")
	for i, paket := range user.Pakets {
		utils.EnterToContinue()
		fmt.Printf("%d. No Resi: %s | Tipe: %s | Dari %s ke %s | Status Terakhir: %s\n",
			i+1, paket.NoResi, paket.Tipe, paket.SenderCity, paket.ReceiverCity,
			paket.Status[len(paket.Status)-1],
		)
	}

}
