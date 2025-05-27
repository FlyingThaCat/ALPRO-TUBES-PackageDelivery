package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
)

func TambahKurir() {
	utils.ClearScreen()
	println("=== Tambah Kurir ===")
	
	var kurir types.User

	kurir.Nama = utils.GetString("Masukkan Nama: ", "Nama tidak boleh kosong.")
	kurir.Username = utils.GetString("Masukkan Username: ", "Username tidak boleh kosong.")
	kurir.Password = utils.GetString("Masukkan Password: ", "Password tidak boleh kosong.")
	kurir.Role = "kurir"
	
	if kurir.Username == "" || kurir.Password == "" {
		fmt.Println("Username dan password tidak boleh kosong.")
		return
	}

	if cariKurir := utils.FindUserByUsername(kurir.Username); cariKurir.Username != "" {
		fmt.Println("Username sudah terdaftar.")
		return
	}

	// Simpan kurir ke database (simulasi)
	datas.UsersDB = append(datas.UsersDB, kurir)
	fmt.Println("Kurir berhasil ditambahkan.")

	utils.EnterToContinue()
}