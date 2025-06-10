package kurir

import (
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"bufio"
	"fmt"
	"os"
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

	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Pilih urutan tampilan paket berdasarkan jarak:")
    fmt.Println("1. Terdekat (ascending)")
    fmt.Println("2. Terjauh (descending)")
    fmt.Print("Masukkan pilihan (1/2): ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    var order string
    switch input {
    case "2":
        order = "desc"
    default:
        order = "asc"
    }

	if order == "asc" {
		sortPaketByStatusAndDistance(pakets, currentCity, "status", order)
	} else {
		sortPaketByStatusAndDistance(pakets, currentCity, "jarak", order)
	}

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

func sortPaketByStatusAndDistance(pakets []types.Paket, currentCity types.Cities, by, order string) {
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
	getDist := func(p types.Paket, status string) float64 {
		city := getReferenceCity(p, status)
		return utils.GetJarak(currentCity, city)
	}

	n := len(pakets)
	asc := strings.ToLower(order) != "desc"

	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			var less bool

			if by == "status" {
				statusI, statusJ := getLastStatus(pakets[idx]), getLastStatus(pakets[j])
				priI, priJ := statusPriority[statusI], statusPriority[statusJ]
				distI, distJ := getDist(pakets[idx], statusI), getDist(pakets[j], statusJ)
				if priJ < priI || (priJ == priI && ((asc && distJ < distI) || (!asc && distJ > distI))) {
					less = true
				}
			} else if by == "jarak" {
				statusIdx := getLastStatus(pakets[idx])
				statusJ := getLastStatus(pakets[j])
				distI := getDist(pakets[idx], statusIdx)
				distJ := getDist(pakets[j], statusJ)
				if asc {
					if distJ < distI {
						less = true
					}
				} else {
					if distJ > distI {
						less = true
					}
				}
			}
			if less {
				idx = j
			}
		}
		if idx != i {
			pakets[i], pakets[idx] = pakets[idx], pakets[i]
		}
	}
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
    // compute distances and store in a slice of pairs
    type entry struct {
        paket types.Paket
        jarak float64
    }
    entries := make([]entry, 0, len(pakets))
    for _, p := range pakets {
        status := p.Status[len(p.Status)-1]
        refCity := getReferenceCity(p, status)
        dist := utils.GetJarak(currentCity, refCity)
        entries = append(entries, entry{paket: p, jarak: dist})
    }

    // print header
    fmt.Println("========================================")
    fmt.Println("📦  DAFTAR PAKET ANDA")
    fmt.Println("========================================")

    // print entries in order
    for i, e := range entries {
        p := e.paket
        status := p.Status[len(p.Status)-1]
        statusKeterangan := getStatusKeterangan(status)
        fmt.Printf(
            "%d. No Resi: %s | Tipe: %s | Dari %s ke %s | Status: %s (%s) | Jarak: %.2f km | Pengirim: %s | Nama Penerima: %s\n",
            i+1, p.NoResi, p.Tipe, p.SenderCity, p.ReceiverCity,
            status, statusKeterangan, e.jarak, p.CreatedBy, p.ReceiverName,
        )
    }

    fmt.Println("========================================")
}