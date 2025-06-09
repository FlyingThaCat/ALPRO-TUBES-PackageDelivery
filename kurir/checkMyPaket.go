package kurir

import (
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getStatusKeterangan(status string) string {
	switch status {
	case "Paket Dibuat":
		return "📦 Perlu Diambil oleh Anda"
	case "Diambil":
		return "✅ Telah Diambil"
	case "Dalam Pengiriman":
		return "🚚 Dalam Pengiriman"
	case "Terkirim":
		return "✔️ Telah Terkirim"
	case "Pengiriman Gagal":
		return "❌ Pengiriman Gagal"
	default:
		return status
	}
}

func CheckMyPaketSorted() {
	utils.ClearScreen()
	username := utils.GetLoggedInUsername()
	pakets := utils.FindPaketByUsername(username)

	if len(pakets) == 0 {
		printNoPackage()
		utils.EnterToContinue()
		return
	}

	currentCity := selectCurrentCity()

	sortPaketByStatusAndDistance(pakets, currentCity)

	printSortedPaketList(pakets, currentCity)
	utils.EnterToContinue()
}

func printNoPackage() {
	fmt.Println("========================================")
	fmt.Println("📦  DAFTAR PAKET ANDA")
	fmt.Println("========================================")
	fmt.Print("\nTidak ada paket yang sedang Anda tangani.\n")
}

func selectCurrentCity() types.Cities {
	fmt.Println("Pilih posisi Anda saat ini:")
	for i, kota := range types.AllCities() {
		fmt.Printf("%d. %s\n", i+1, kota)
	}
	fmt.Print("Masukkan nomor kota Anda: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(types.Kota) {
		fmt.Println("Pilihan tidak valid. Gunakan default Jakarta.")
		choice = 1
	}
	return types.ParseCity(types.Kota[choice-1])
}

func sortPaketByStatusAndDistance(pakets []types.Paket, currentCity types.Cities) {
	statusPriority := map[string]int{
		"Paket Dibuat":     1,
		"Diambil":          2,
		"Dalam Pengiriman": 3,
		"Terkirim":         4,
		"Pengiriman Gagal": 5,
	}

	getLastStatus := func(p types.Paket) string {
		if len(p.Status) == 0 {
			return ""
		}
		return p.Status[len(p.Status)-1]
	}

	sort.SliceStable(pakets, func(i, j int) bool {
		statusI := getLastStatus(pakets[i])
		statusJ := getLastStatus(pakets[j])

		if statusPriority[statusI] != statusPriority[statusJ] {
			return statusPriority[statusI] < statusPriority[statusJ]
		}

		cityI := getReferenceCity(pakets[i], statusI)
		cityJ := getReferenceCity(pakets[j], statusJ)

		distI := utils.GetJarak(currentCity, cityI)
		distJ := utils.GetJarak(currentCity, cityJ)

		return distI < distJ
	})
}

func getReferenceCity(paket types.Paket, status string) types.Cities {
	switch status {
	case "Dalam Pengiriman":
		return paket.ReceiverCity
	case "Diambil":
		return paket.ReceiverCity
	default:
		return paket.SenderCity
	}
}

func printSortedPaketList(pakets []types.Paket, currentCity types.Cities) {
	fmt.Println("========================================")
	fmt.Println("📦  DAFTAR PAKET ANDA")
	fmt.Println("========================================")

	for i, paket := range pakets {
		status := paket.Status[len(paket.Status)-1]
		statusKeterangan := getStatusKeterangan(status)
		refCity := getReferenceCity(paket, status)
		jarak := utils.GetJarak(currentCity, refCity)

		fmt.Printf("%d. No Resi: %s | Tipe: %s | Dari %s ke %s | Status: %s (%s) | Jarak: %.2f km\n",
			i+1, paket.NoResi, paket.Tipe, paket.SenderCity, paket.ReceiverCity,
			status, statusKeterangan, jarak,
		)
	}
	fmt.Println("========================================")
}
