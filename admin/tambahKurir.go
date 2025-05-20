package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
)

func TambahKurir() {
	utils.ClearScreen()
	// Fungsi untuk menambahkan kurir baru
	var kurir types.User

	println("=== Tambah Kurir ===")
	print("Masukkan Nama: ")
	fmt.Scanln(&kurir.Nama)
	print("Masukkan Username: ")
	fmt.Scanln(&kurir.Username)
	print("Masukkan Password: ")
	fmt.Scanln(&kurir.Password)

	kurir.Role = "kurir"
	
	// Validasi username dan password
	if kurir.Username == "" || kurir.Password == "" {
		fmt.Println("Username dan password tidak boleh kosong.")
		return
	}

	if cariKurir := datas.FindUserByUsername(kurir.Username); cariKurir.Username != "" {
		fmt.Println("Username sudah terdaftar.")
		return
	}

	// Simpan kurir ke database (simulasi)
	datas.UsersDB = append(datas.UsersDB, kurir)
	fmt.Println("Kurir berhasil ditambahkan.")

	utils.EnterToContinue()
}