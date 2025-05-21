package kurir

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
	"strings"
	"time"
)

func filterPaketByKurir(username string) []types.Paket {
	var result []types.Paket
	for _, p := range datas.PaketDB {
		if p.Kurir == username {
			result = append(result, p)
		}
	}
	return result
}

// updateStatusPaket cari paket berdasarkan NoResi lalu update status terakhirnya
func updateStatusPaket(noResi, statusBaru string) bool {
	for i := range datas.PaketDB {
		if datas.PaketDB[i].NoResi == noResi {
			datas.PaketDB[i].Status = append(datas.PaketDB[i].Status, statusBaru)
			datas.PaketDB[i].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

// UpdateStatus menampilkan paket yang ditangani kurir dan mengupdate status paket yang dipilih
func UpdateStatus() {
	utils.ClearScreen()

	username := utils.GetLoggedInUsername()

	paketList := filterPaketByKurir(username)
	if len(paketList) == 0 {
		fmt.Println("Tidak ada paket yang sedang Anda tangani.")
		utils.EnterToContinue()
		return
	}

	fmt.Println("Daftar paket yang Anda tangani:")
	for i, paket := range paketList {
		lastStatus := "Belum ada status"
		if len(paket.Status) > 0 {
			lastStatus = paket.Status[len(paket.Status)-1]
		}
		fmt.Printf("%d. NoResi: %s, Status terakhir: %s\n", i+1, paket.NoResi, lastStatus)
		utils.EnterToContinue()
	}

	fmt.Print("\nMasukkan nomor urut paket yang ingin diupdate statusnya: ")
	utils.EnterToContinue()
	var pilih int
	_, err := fmt.Scanln(&pilih)
	if err != nil || pilih < 1 || pilih > len(paketList) {
		fmt.Println("Pilihan tidak valid.")
		utils.EnterToContinue()
		return
	}

	selectedPaket := &paketList[pilih-1]

	fmt.Println("Masukkan status baru paket (contoh: Diambil, Dalam Pengiriman, Terkirim):")
	var statusBaru string
	fmt.Scanln(&statusBaru)
	statusBaru = strings.TrimSpace(statusBaru)
	if statusBaru == "" {

		fmt.Println("Status tidak boleh kosong.")
		utils.EnterToContinue()
		return
	}

	// Update status paket di DB utama
	if updateStatusPaket(selectedPaket.NoResi, statusBaru) {

		fmt.Println("Status paket berhasil diperbarui.")
		utils.EnterToContinue()

	} else {

		fmt.Println("Gagal memperbarui status paket.")
		utils.EnterToContinue()

	}
}
