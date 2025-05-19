package users

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"fmt"
	"time"
)

func TambahPaket() {
	var input types.Paket
	var confirm string

	fmt.Print("\n=== Tambah Paket ===\n")
	fmt.Println("Tipe Paket:")
	for _, tipe := range types.PackageTypes {
		fmt.Printf("- %s\n", tipe)
	}

	fmt.Print("Masukkan tipe paket: ")
	fmt.Scanln(&input.Tipe)

	if !types.IsValidPackageType(input.Tipe) {
		fmt.Println("Tipe paket tidak valid.")
		return
	}

	fmt.Print("Masukkan berat paket (kg): ")
	fmt.Scanln(&input.Berat)

	fmt.Println("Kota Pengirim: ")
	for _, tipe := range types.Kota {
		fmt.Printf("- %s\n", tipe)
	}
	fmt.Print("Masukkan kota pengirim: ")
	fmt.Scanln(&input.SenderCity)

	fmt.Println("Kota Tujuan:")
	for _, tipe := range types.Kota {
		fmt.Printf("- %s\n", tipe)
	}
	fmt.Print("Masukkan kota tujuan: ")
	fmt.Scanln(&input.ReceiverCity)

	// calculate price (for now dummy value)
	input.Harga = 10000.0

	fmt.Printf("\n=== Konfirmasi Paket ===\n")
	fmt.Println("Harga paket: Rp", input.Harga)
	fmt.Printf("Apakah Anda yakin ingin menambahkan paket ini? (y/n): ")
	fmt.Scanln(&confirm)
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Paket tidak ditambahkan.")
		return
	}

	// generate unique tracking number
	input.NoResi = fmt.Sprintf("%d", len(datas.PaketDB)+1)
	input.CreatedAt = time.Now()
	// when creating a new package, the status is always "Dibuat"
	input.Status = []string{"Paket Dibuat"}

	datas.PaketDB = append(datas.PaketDB, input)
	fmt.Println("Paket berhasil ditambahkan.")

	fmt.Println("=== Detail Paket ===")
	fmt.Println("No Resi:", datas.PaketDB[len(datas.PaketDB)-1].NoResi)
	fmt.Println("Tipe:", input.Tipe)
	fmt.Println("Berat:", input.Berat)
	fmt.Println("Harga:", input.Harga)
	fmt.Println("Kota Pengirim:", input.SenderCity)
	fmt.Println("Kota Tujuan:", input.ReceiverCity)
	fmt.Println("Status:", input.Status)
	fmt.Println("Dibuat pada:", input.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("Paket berhasil ditambahkan.")
	fmt.Println("=== Detail Paket ===")
}

func MenuUser() {
	fmt.Println("=== Menu User ===")
	fmt.Println("1. Tambah Paket")
	fmt.Println("2. Cek Paket")
	fmt.Println("3. Logout")
}