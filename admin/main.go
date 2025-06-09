package admin

import (
	"PackageDelivery/utils"
	"fmt"
	"strings"
)

func MenuAdmin() {
	utils.ClearScreen()

	username := utils.GetLoggedInUsername()

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("🛠️  MENU ADMIN, Hello", username)
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("📦  Paket")
	fmt.Println("1️⃣  Tambah Paket Baru")
	fmt.Println("2️⃣  Lihat Semua Paket")
	fmt.Println("3️⃣  Edit Paket")
	fmt.Println("4️⃣  Hapus Paket")

	fmt.Println("\n🚚  Kurir")
	fmt.Println("5️⃣  Tambah Kurir Baru")
	fmt.Println("6️⃣  Lihat Semua Kurir")
	fmt.Println("7️⃣  Edit Kurir")
	fmt.Println("8️⃣  Hapus Kurir")
	fmt.Println("9️⃣  Assign Paket ke Kurir")

	//
	fmt.Println("\n0️⃣  Logout")
	fmt.Println(strings.Repeat("=", 40))
}
