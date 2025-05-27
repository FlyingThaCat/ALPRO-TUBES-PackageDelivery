package admin

import (
	"PackageDelivery/utils"
	"fmt"
)

func HapusKurir() {
	utils.ClearScreen()
	fmt.Println("=== Hapus Kurir ===")

	username := utils.GetString("Masukkan username kurir yang ingin dihapus: ", "Silakan masukkan username kurir yang valid.")
	
	ok := utils.DeleteUser(username)
	
	if !ok {
		fmt.Println("Kurir tidak dapat dihapus. Username tidak ditemukan.")
		return
	}

	fmt.Printf("Kurir dengan username %s telah dihapus.\n\n", username)

	utils.EnterToContinue()
}