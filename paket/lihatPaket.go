package paket

import (
	"PackageDelivery/datas"
	"fmt"
)

func LihatPaket() {
	fmt.Printf("\n=== Daftar Paket ===\n")
	fmt.Println("No Resi\t | Status Terakhir\t\t\tKurir")
	for _, paket := range datas.PaketDB {
		kurir := paket.Kurir
		if kurir == "" {
			kurir = "Belum Ditugaskan"
		}
		fmt.Printf("%s\t | %s\t\t\t\t%s\n", paket.NoResi, paket.Status[len(paket.Status)-1], kurir)
	}
	fmt.Printf("\n====================\n")
}