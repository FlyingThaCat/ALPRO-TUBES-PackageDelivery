package paket

import (
	"PackageDelivery/admin"
	"PackageDelivery/utils"
	"fmt"
)

func AssignPaket() {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("📦  ASSIGN PAKET KE KURIR")
	fmt.Println("========================================")

	LihatPaket(false, true)
	admin.LihatKurir(false)

	fmt.Println("----------------------------------------")
	noResi := utils.GetString("Masukkan No Resi Paket yang ingin di-assign: ", "Silakan masukkan No Resi yang valid.")
	username := utils.GetString("Masukkan Username Kurir yang akan mengantarkan paket: ", "Silakan masukkan username kurir yang valid.")
	foundPaket := utils.FindPaketByNoResi(noResi)

	if foundPaket.Kurir != "" {
		fmt.Println("❌ Paket sduah ditugaskan ke ", foundPaket.Kurir)
		utils.EnterToContinue()
		return
	}

	if foundPaket.NoResi == "" {
		fmt.Println("❌ Paket tidak ditemukan.")
		utils.EnterToContinue()
		return
	}

	foundKurir := utils.FindUserByUsername(username)
	if foundKurir.Username == "" {
		fmt.Println("❌ Kurir tidak ditemukan.")
		utils.EnterToContinue()
		return
	}
	if foundKurir.Role != "kurir" {
		fmt.Println("❌ Username yang dimasukkan bukan kurir.")
		utils.EnterToContinue()
		return
	}

	if utils.AssignPaketToKurir(foundKurir.Username, foundPaket.NoResi) {
		fmt.Println("✅ Paket berhasil ditugaskan ke kurir.")
	} else {
		fmt.Println("❌ Gagal menugaskan paket.")
	}

	fmt.Println("========================================")
	utils.EnterToContinue()
}
