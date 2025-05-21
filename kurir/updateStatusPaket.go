package kurir

import (
	"PackageDelivery/utils"
	"fmt"
)

func UpdateStatus() {
	username := utils.GetLoggedInUsername()

	fmt.Printf("Selamat datang kembali, kurir %s!\n", username)

}
