package paket

import (
	"PackageDelivery/datas"
	"PackageDelivery/utils"
	"fmt"
	"os"
	"text/tabwriter"
)

func LihatPaket(clear bool) {
	if clear {
		utils.ClearScreen()
	}
	fmt.Println("\n=== Daftar Paket ===")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "No Resi\tTipe\tBerat\tHarga\tKota Pengirim\tKota Tujuan\tStatus Terakhir\tDibuat Pada\tDiperbarui Pada\tKurir")

	for _, paket := range datas.PaketDB {
		kurir := paket.Kurir
		if kurir == "" {
			kurir = "Belum Ditugaskan"
		}

		// Cek status terakhir agar tidak panic jika kosong
		statusTerakhir := "-"
		if len(paket.Status) > 0 {
			statusTerakhir = paket.Status[len(paket.Status)-1]
		}

		fmt.Fprintf(w, "%s\t%s\t%.1f\t%.0f\t%s\t%s\t%s\t%s\t%s\t%s\n",
			paket.NoResi,
			paket.Tipe,
			paket.Berat,
			paket.Harga,
			paket.SenderCity,
			paket.ReceiverCity,
			statusTerakhir,
			paket.CreatedAt.Format("2006-01-02 15:04:05"),
			paket.UpdatedAt.Format("2006-01-02 15:04:05"),
			kurir,
		)
	}

	w.Flush()
	fmt.Print("====================\n\n")

	if clear {
		utils.EnterToContinue()
	}
}
