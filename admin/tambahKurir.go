package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
)

func TambahKurir() {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("🚚  TAMBAH KURIR")
	fmt.Println("========================================")

	var kurir types.User

	kurir.Nama = utils.GetString("Masukkan Nama: ", "Nama tidak boleh kosong.")
	kurir.Username = utils.GetString("Masukkan Username: ", "Username tidak boleh kosong.")
	kurir.Password = utils.GetString("Masukkan Password: ", "Password tidak boleh kosong.")
	kurir.Role = "kurir"

	if kurir.Username == "" || kurir.Password == "" {
		fmt.Println("\n⚠️  Username dan password tidak boleh kosong.")
		utils.EnterToContinue()
		return
	}

	if cariKurir := utils.FindUserByUsername(kurir.Username); cariKurir.Username != "" {
		fmt.Println("\n⚠️  Username sudah terdaftar.")
		utils.EnterToContinue()
		return
	}

	// Simpan kurir ke database (simulasi)
	datas.UsersDB = append(datas.UsersDB, kurir)

	fmt.Println("\n✅  Kurir berhasil ditambahkan.")
	fmt.Println("========================================")

	utils.EnterToContinue()
}
