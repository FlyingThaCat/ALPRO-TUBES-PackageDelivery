package kurir

import (
	"PackageDelivery/utils"
	"fmt"
)

func MenuKurir() {
	utils.ClearScreen()

	fmt.Println("=== Menu Kurir ===")
	fmt.Println("1. Lihat Assigned Paket")
	fmt.Println("2. Update Status Paket")
	fmt.Println("0. Logout")
}
