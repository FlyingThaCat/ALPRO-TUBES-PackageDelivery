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

	found := utils.FindUserByUsername(username)
	if found.Username == "" {
		fmt.Println("\n⚠️  Kurir tidak ditemukan.")
		utils.EnterToContinue()
		return
	}

	pakets := utils.FindPaketByUsername(username)

	if len(pakets) > 0 {
		fmt.Println("\n⚠️  Kurir tidak dapat dihapus. Masih ada paket yang ditugaskan ke kurir ini.")
		utils.EnterToContinue()
		return
	}

	if !utils.GetConfirmation("Apakah Anda yakin ingin menghapus kurir ini?", "Silakan masukkan 'y' untuk ya atau 'n' untuk tidak.") {
		fmt.Println("Kurir Tidak Jadi Dihapus.")
		utils.EnterToContinue()
		return
	}

	ok := utils.DeleteUser(username)

	if !ok {
		fmt.Println("\n⚠️  Kurir tidak dapat dihapus. Mungkin ada kesalahan pada sistem.")
		fmt.Println("⚠️  Pastikan tidak ada paket yang ditugaskan ke kurir ini.")
		utils.EnterToContinue()
		return
	}

	fmt.Printf("\n✅ Kurir dengan username '%s' telah dihapus.\n\n", username)
	utils.EnterToContinue()
}
