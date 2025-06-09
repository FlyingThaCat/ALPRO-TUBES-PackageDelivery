package admin

import (
	"PackageDelivery/utils"
	"fmt"
)

func HapusKurir() {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("🗑️  HAPUS KURIR")
	fmt.Println("========================================")

	username := utils.GetString("Masukkan username kurir yang ingin dihapus: ", "Silakan masukkan username kurir yang valid.")

	ok := utils.DeleteUser(username)

	if !ok {
		fmt.Println("\n⚠️  Kurir tidak dapat dihapus. Username tidak ditemukan.")
		utils.EnterToContinue()
		return
	}

	fmt.Printf("\n✅ Kurir dengan username '%s' telah dihapus.\n\n", username)
	utils.EnterToContinue()
}
