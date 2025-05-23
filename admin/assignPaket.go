package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/paket"
	"PackageDelivery/utils"
	"fmt"
)

func AssignPaket() {
	utils.ClearScreen()
	
	var noResi string
	var username string
	paket.LihatPaket(false, true)
	LihatKurir(false)
	println("=== Assign Paket ===")
	print("Masukkan No Resi Paket yang ingin diassign: ")
	fmt.Scanln(&noResi)
	print("Masukkan Username Kurir yang akan mengantarkan paket: ")
	fmt.Scanln(&username)

	foundPaket := datas.CariPaket(noResi)
	if foundPaket.NoResi == "" {
		fmt.Println("Paket tidak ditemukan.")
		return
	}
	foundKurir := datas.FindUserByUsername(username)
	if foundKurir.Username == "" {
		fmt.Println("Kurir tidak ditemukan.")
		return
	}
	if foundKurir.Role != "kurir" {
		fmt.Println("Username yang dimasukkan bukan kurir.")
		return
	}

	if datas.AssignPaketToKurir(foundKurir.Username, foundPaket.NoResi) {
		fmt.Println("Paket berhasil diassign ke kurir.")
	} else {
		fmt.Println("Paket gagal diassign.")
	}
	
	utils.EnterToContinue()
}