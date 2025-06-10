package paket

import (
	"PackageDelivery/datas"
	"PackageDelivery/utils"
	"fmt"
	"os"
	"text/tabwriter"
)

func LihatPaket(clear bool, hideAssigned bool) {
	if clear {
		utils.ClearScreen()
	}

	fmt.Println("========================================")
	fmt.Println("📦  DAFTAR PAKET")
	fmt.Println("========================================")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "No Resi\tTipe\tBerat (kg)\tHarga (Rp)\tKota Pengirim\tKota Tujuan\tStatus Terakhir\tDibuat Pada\tDiperbarui Pada\tKurir\tDibuat Oleh")

	for _, paket := range datas.PaketDB {
		if hideAssigned && paket.Kurir != "" {
			continue
		}

		kurir := paket.Kurir
		if kurir == "" {
			kurir = "Belum Ditugaskan"
		}

		statusTerakhir := "-"
		if len(paket.Status) > 0 {
			statusTerakhir = paket.Status[len(paket.Status)-1]
		}

		fmt.Fprintf(w, "%s\t%s\t%.1f\t%.0f\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
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
			paket.CreatedBy,
		)
	}

	w.Flush()

	if clear {
		utils.EnterToContinue()
	}
}
