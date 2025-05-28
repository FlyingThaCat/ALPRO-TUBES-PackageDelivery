package paket

import (
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
)

func DetailPaket(paket types.Paket){
	fmt.Print("\n\n=== Detail Paket ===\n")
	fmt.Println("No Resi:", paket.NoResi)
	fmt.Println("Tipe:", paket.Tipe)
	fmt.Println("Berat:", paket.Berat)
	fmt.Printf("Harga: %.0f \n", paket.Harga)
	fmt.Println("Kota Pengirim:", paket.SenderCity)
	fmt.Println("Kota Tujuan:", paket.ReceiverCity)
	fmt.Println("Status:", paket.Status)
	fmt.Println("Dibuat pada:", paket.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("Diperbarui pada:", paket.UpdatedAt.Format("2006-01-02 15:04:05"))
	if paket.Kurir != "" {
		fmt.Println("Kurir:", paket.Kurir)
	} else {
		fmt.Println("Kurir: Belum Ditugaskan")
	}
	fmt.Println("Paket berhasil ditambahkan.")
	fmt.Print("=== Detail Paket ===\n\n")

	utils.EnterToContinue()
}