package kurir

import (
	"PackageDelivery/utils"
	"fmt"
)

func UpdateStatus() {
	username := utils.GetLoggedInUsername()

	fmt.Printf("Selamat datang kembali, kurir %s!\n", username)

	// user := datas.FindUserByUsername(username)
	// if user.Username == "" {
	// 	fmt.Println("User tidak ditemukan.")
	// 	return
	// }

	// if user.Role != "kurir" {
	// 	fmt.Println("Hanya kurir yang dapat mengakses fitur ini.")
	// 	return
	// }

	// if len(user.Pakets) == 0 {
	// 	fmt.Println("Tidak ada paket yang sedang Anda tangani.")
	// 	return
	// }

	// fmt.Println("=== Daftar Paket Anda ===")
	// for i, paket := range user.Pakets {
	// 	fmt.Printf("%d. No Resi: %s | Tipe: %s | Dari %s ke %s | Status Terakhir: %s\n",
	// 		i+1, paket.NoResi, paket.Tipe, paket.SenderCity, paket.ReceiverCity,
	// 		paket.Status[len(paket.Status)-1],
	// 	)
	// }
}
