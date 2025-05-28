package kurir

import (
	"PackageDelivery/utils"
	"fmt"
	"strings"
)

func MenuKurir() {
	utils.ClearScreen()

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("🚚  MENU KURIR")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("1️⃣  Lihat Assigned Paket")
	fmt.Println("2️⃣  Update Status Paket")
	fmt.Println("0️⃣  Logout")
	fmt.Println(strings.Repeat("=", 40))
}
