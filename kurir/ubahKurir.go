package kurir

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)


func UbahKurir(kurir types.User) {
	utils.ClearScreen()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Ubah Kurir ===")
	
	inputWithDefault := func(prompt string, current string) string {
		fmt.Printf("%s [%s]: ", prompt, current)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			return current
		}
		return text
	}

	oldUsername := kurir.Username

	kurir.Nama = inputWithDefault("Nama", string(kurir.Nama))
	kurir.Username = inputWithDefault("Username", string(kurir.Username))
	kurir.Password = inputWithDefault("Password", string(kurir.Password))
	
	// Simpan perubahan ke database (simulasi)
	for i, p := range datas.UsersDB {
		if p.Username == oldUsername {
			datas.UsersDB[i] = kurir
			break
		}
	}

	utils.ClearScreen()
	fmt.Println("Kurir berhasil diubah:")
	fmt.Printf("Nama: %s\n", kurir.Nama)
	fmt.Printf("Username: %s\n", kurir.Username)
	fmt.Printf("Password: %s\n", kurir.Password)
	utils.EnterToContinue()
}
