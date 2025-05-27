package paket

import (
	"PackageDelivery/admin"
	"PackageDelivery/utils"
	"fmt"
)

func AssignPaket() {
	utils.ClearScreen()
	println("=== Assign Paket ===")
	
	LihatPaket(false, true)
	admin.LihatKurir(false)
	
	noResi := utils.GetString("Masukkan No Resi Paket yang ingin diassign: ", "Silakan masukkan No Resi yang valid.")
	username := utils.GetString("Masukkan Username Kurir yang akan mengantarkan paket: ", "Silakan masukkan username kurir yang valid.")

	foundPaket := utils.FindPaketByNoResi(noResi)
	if foundPaket.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
		return
	}
	foundKurir := utils.FindUserByUsername(username)
	if foundKurir.Username == "" {
		fmt.Println("Kurir tidak ditemukan.")
		return
	}
	if foundKurir.Role != "kurir" {
		fmt.Println("Username yang dimasukkan bukan kurir.")
		return
	}

	if utils.AssignPaketToKurir(foundKurir.Username, foundPaket.NoResi) {
		fmt.Println("Paket berhasil diassign ke kurir.")
	} else {
		fmt.Println("Paket gagal diassign.")
	}
	
	utils.EnterToContinue()
}