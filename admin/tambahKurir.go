package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"fmt"
)

func TambahKurir() {
	// Fungsi untuk menambahkan kurir baru
	var username, password string

	println("=== Tambah Kurir ===")
	print("Masukkan Username: ")
	fmt.Scanln(&username)
	print("Masukkan Password: ")
	fmt.Scanln(&password)

	// Validasi username dan password
	if username == "" || password == "" {
		fmt.Println("Username dan password tidak boleh kosong.")
		return
	}

	// Simpan kurir ke database (simulasi)
	datas.UsersDB = append(datas.UsersDB, types.User{Username: username, Password: password, Role: "kurir"})
	fmt.Println("Kurir berhasil ditambahkan.")
}