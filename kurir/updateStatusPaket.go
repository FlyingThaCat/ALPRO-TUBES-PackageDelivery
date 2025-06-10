package kurir

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
	"strings"
	"time"
)

var statusPriority = map[string]int{
	"Paket Dibuat":     1,
	"Diambil":          2,
	"Dalam Pengiriman": 3,
	"Terkirim":         4,
	"Pengiriman Gagal": 5,
}

func filterPaketByKurir(username string) []types.Paket {
	var result []types.Paket
	for _, p := range datas.PaketDB {
		if p.Kurir == username {
			result = append(result, p)
		}
	}
	return result
}

func updateStatusPaket(noResi, statusBaru string) bool {
	for i := range datas.PaketDB {
		if datas.PaketDB[i].NoResi == noResi {
			lastStatus := ""
			if len(datas.PaketDB[i].Status) > 0 {
				lastStatus = datas.PaketDB[i].Status[len(datas.PaketDB[i].Status)-1]
			}
			if lastStatus == statusBaru {
				return false
			}
			if statusPriority[lastStatus] > statusPriority[statusBaru] && statusBaru != "Pengiriman Gagal" {
				return false
			}

			datas.PaketDB[i].Status = append(datas.PaketDB[i].Status, statusBaru)
			datas.PaketDB[i].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

func pilihPaket(paketList []types.Paket) *types.Paket {
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
		utils.EnterToContinue()
		return nil
	}

	return &paketList[pilih-1]
}

func getStatusOptions(lastStatus string) []string {
	options := []string{}
	validStatuses := []string{"Diambil", "Dalam Pengiriman", "Terkirim", "Pengiriman Gagal"}

	for _, s := range validStatuses {
		if s == "Pengiriman Gagal" || statusPriority[s] > statusPriority[lastStatus] {
			options = append(options, s)
		}
	}
	return options
}

func pilihStatus(statusOptions []string) string {
	fmt.Println("\nPilih status baru untuk paket:")
	for i, s := range statusOptions {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	fmt.Print("Masukkan nomor status baru: ")
	var statusChoice int
	_, err := fmt.Scanln(&statusChoice)
	if err != nil || statusChoice < 1 || statusChoice > len(statusOptions) {
		fmt.Println("Pilihan status tidak valid.")
		utils.EnterToContinue()
		return ""
	}

	return statusOptions[statusChoice-1]
}

func konfirmasiUpdate(noResi, statusBaru string) bool {
	fmt.Printf("Anda akan mengubah status paket NoResi %s menjadi \"%s\". Lanjutkan? (y/n): ", noResi, statusBaru)
	var confirm string
	fmt.Scanln(&confirm)
	confirm = strings.TrimSpace(strings.ToLower(confirm))
	return confirm == "y"
}

func UpdateStatus() {
	utils.ClearScreen()

	username := utils.GetLoggedInUsername()
	paketList := filterPaketByKurir(username)
	if len(paketList) == 0 {
		fmt.Println("Tidak ada paket yang sedang Anda tangani.")
		utils.EnterToContinue()
		return
	}

	selectedPaket := pilihPaket(paketList)
	if selectedPaket == nil {
		return
	}

	lastStatus := ""
	if len(selectedPaket.Status) > 0 {
		lastStatus = selectedPaket.Status[len(selectedPaket.Status)-1]
	}

	statusOptions := getStatusOptions(lastStatus)
	if len(statusOptions) == 0 {
		fmt.Println("Tidak ada status baru yang bisa dipilih berdasarkan status terakhir.")
		utils.EnterToContinue()
		return
	}

	statusBaru := pilihStatus(statusOptions)
	if statusBaru == "" {
		return
	}

	if !konfirmasiUpdate(selectedPaket.NoResi, statusBaru) {
		fmt.Println("Update status dibatalkan.")
		utils.EnterToContinue()
		return
	}

	if updateStatusPaket(selectedPaket.NoResi, statusBaru) {
		fmt.Println("Status paket berhasil diperbarui.")
	} else {
		fmt.Println("Gagal memperbarui status paket. Pastikan status baru valid dan tidak mundur.")
	}

	utils.EnterToContinue()
}
