package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
	"os"
	"text/tabwriter"
)

func LihatKurir(clear bool) {
	if clear {
		utils.ClearScreen()
	}

	fmt.Println("========================================")
	fmt.Println("🚚  DAFTAR KURIR")
	fmt.Println("========================================")

	var kurirs []types.User
	for _, user := range datas.UsersDB {
		if user.Role == "kurir" {
			kurirs = append(kurirs, user)
		}
	}

	// Sorting kurir berdasarkan jumlah paket yang diassign, ascending
	for i := 1; i < len(kurirs); i++ {
		temp := kurirs[i]
		j := i - 1
		for j >= 0 && len(kurirs[j].Pakets) > len(temp.Pakets) {
			kurirs[j+1] = kurirs[j]
			j--
		}
		kurirs[j+1] = temp
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Nama\tUsername\tPassword\tPaket yang Diassign")
	for _, kurir := range kurirs {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\n", kurir.Nama, kurir.Username, kurir.Password, len(kurir.Pakets))
	}
	w.Flush()

	fmt.Println("========================================")
	if clear {
		utils.EnterToContinue()
	}
}
