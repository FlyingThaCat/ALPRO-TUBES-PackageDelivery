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
	fmt.Println("\n=== Daftar Kurir ===")	

	var kurirs []types.User
	for _, user := range datas.UsersDB {
		if user.Role == "kurir" {
			kurirs = append(kurirs, user)
		}
	}

	var paketCounts = make([]int, len(kurirs))
	for i, k := range kurirs {
		paketCounts[i] = len(k.Pakets)
	}

	
	var pairs = make([]types.RowPair, len(paketCounts))
	for i := range paketCounts {
		pairs[i] = types.RowPair{Index: i, Count: paketCounts[i]}
	}

	var indexes = make([]int, len(pairs))
	for i := range pairs {
		indexes[i] = pairs[i].Count
	}
	utils.SortArray(indexes, "asc")

	var sortedKurirs []types.User
	for _, sortedCount := range indexes {
		for i, p := range pairs {
			if p.Count == sortedCount && p.Index != -1 {
				sortedKurirs = append(sortedKurirs, kurirs[p.Index])
				pairs[i].Index = -1
				break
			}
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Nama\tUsername\tPassword\tPaket yang Diassign")
	for _, kurir := range sortedKurirs {
		if kurir.Role != "kurir" {
			continue
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n", kurir.Nama, kurir.Username, kurir.Password, len(kurir.Pakets))
	}
	w.Flush()
	fmt.Print("====================\n\n")
	if clear {
		utils.EnterToContinue()
	}
}