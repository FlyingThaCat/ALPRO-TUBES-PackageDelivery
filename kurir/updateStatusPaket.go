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
	}

	fmt.Print("\nMasukkan nomor urut paket yang ingin diupdate statusnya: ")
	var pilih int
	_, err := fmt.Scanln(&pilih)
	if err != nil || pilih < 1 || pilih > len(paketList) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	selectedPaket := &paketList[pilih-1]

	// Daftar status enum
	statusOptions := []string{
		"Diambil",
		"Dalam Pengiriman",
		"Terkirim",
		"Pengiriman Gagal",
	}

	fmt.Println("\nPilih status baru untuk paket:")
	for i, s := range statusOptions {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	fmt.Print("Masukkan nomor status baru: ")
	var statusChoice int
	_, err = fmt.Scanln(&statusChoice)
	if err != nil || statusChoice < 1 || statusChoice > len(statusOptions) {
		fmt.Println("Pilihan status tidak valid.")
		return
	}

	statusBaru := statusOptions[statusChoice-1]

	// Konfirmasi
	fmt.Printf("Anda akan mengubah status paket NoResi %s menjadi \"%s\". Lanjutkan? (y/n): ", selectedPaket.NoResi, statusBaru)
	var confirm string
	fmt.Scanln(&confirm)
	confirm = strings.TrimSpace(strings.ToLower(confirm))
	if confirm != "y" {
		fmt.Println("Update status dibatalkan.")
		utils.EnterToContinue()
		return
	}

	// Update status di database utama
	if updateStatusPaket(selectedPaket.NoResi, statusBaru) {
		fmt.Println("Status paket berhasil diperbarui.")
	} else {
		fmt.Println("Gagal memperbarui status paket.")
	}

	utils.EnterToContinue()
}
