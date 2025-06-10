package paket

import (
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
	"strings"
)

func DetailPaket(paket types.Paket) {
	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("📦  DETAIL PAKET")
	fmt.Println("========================================")

	fmt.Println("No Resi       :", paket.NoResi)
	fmt.Println("Tipe          :", paket.Tipe)
	fmt.Printf("Berat         : %.2f kg\n", paket.Berat)
	fmt.Printf("Harga         : Rp %.0f\n", paket.Harga)
	fmt.Println("Kota Pengirim :", paket.SenderCity)
	fmt.Println("Kota Tujuan   :", paket.ReceiverCity)

	// Status
	if len(paket.Status) > 0 {
		fmt.Println("Status        :", strings.Join(paket.Status, ", "))
	} else {
		fmt.Println("Status        : -")
	}

	fmt.Println("Dibuat Pada   :", paket.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("Diperbarui    :", paket.UpdatedAt.Format("2006-01-02 15:04:05"))

	// Kurir
	if paket.Kurir != "" {
		fmt.Println("Kurir         :", paket.Kurir)
	} else {
		fmt.Println("Kurir         : Belum Ditugaskan")
	}

	fmt.Println("========================================")

}
