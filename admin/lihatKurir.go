package admin

import (
	"PackageDelivery/datas"
	"PackageDelivery/utils"
	"fmt"
	"os"
	"text/tabwriter"
)

func LihatKurir() {
	utils.ClearScreen()
	fmt.Println("\n=== Daftar Kurir ===")
	
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Nama\tUsername\tPassword\tPaket")

	for _, kurir := range datas.UsersDB {
		if kurir.Role != "kurir" {
			continue
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n", kurir.Nama, kurir.Username, kurir.Password, kurir.Pakets)
	}
	w.Flush()
	fmt.Print("====================\n\n")
	utils.EnterToContinue()
}