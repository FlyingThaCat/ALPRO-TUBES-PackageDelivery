package paket

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func UbahPaket(paket types.Paket) {
	utils.ClearScreen()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========================================")
	fmt.Println("✏️  UBAH DATA PAKET")
	fmt.Println("========================================")

	inputWithDefault := func(prompt string, current string) string {
		fmt.Printf("%s [%s]: ", prompt, current)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			return current
		}
		return text
	}

	paket.Tipe = inputWithDefault("Tipe Paket", string(paket.Tipe))

	beratStr := inputWithDefault("Berat (kg)", fmt.Sprintf("%.2f", paket.Berat))
	fmt.Sscanf(beratStr, "%f", &paket.Berat)

	// Ambil kota pengirim dan validasi
	senderCityStr := inputWithDefault("Kota Pengirim", string(paket.SenderCity))
	senderCity, err := utils.ParseCity(senderCityStr)
	if err != nil {
		fmt.Printf("\n⚠️  Kota pengirim tidak valid: %s\n", err)
		utils.EnterToContinue()
		return
	}
	paket.SenderCity = senderCity

	// Ambil kota tujuan dan validasi
	receiverCityStr := inputWithDefault("Kota Tujuan", string(paket.ReceiverCity))
	receiverCity, err := utils.ParseCity(receiverCityStr)
	if err != nil {
		fmt.Printf("\n⚠️  Kota tujuan tidak valid: %s\n", err)
		utils.EnterToContinue()
		return
	}
	paket.ReceiverCity = receiverCity

	// Ambil status
	statusStr := inputWithDefault("Status Paket (pisahkan dengan koma)", strings.Join(paket.Status, ","))
	paket.Status = strings.Split(statusStr, ",")

	// Hitung ulang harga berdasarkan koordinat dan data paket
	senderLat, senderLon, err := utils.GetCoordinates(paket.SenderCity)
	if err != nil {
		fmt.Printf("\n⚠️ Error koordinat kota pengirim: %s\n", err)
		utils.EnterToContinue()
		return
	}

	receiverLat, receiverLon, err := utils.GetCoordinates(paket.ReceiverCity)
	if err != nil {
		fmt.Printf("\n⚠️ Error koordinat kota tujuan: %s\n", err)
		utils.EnterToContinue()
		return
	}

	_, newPrice := utils.CalculatePackagePrice(senderLat, senderLon, receiverLat, receiverLon, paket.Berat, strings.ToLower(paket.Tipe))
	paket.Harga = newPrice

	paket.UpdatedAt = time.Now()

	// Update paket di database
	for i, p := range datas.PaketDB {
		if p.NoResi == paket.NoResi {
			datas.PaketDB[i] = paket
			break
		}
	}

	utils.ClearScreen()
	fmt.Println("========================================")
	fmt.Println("✅  DATA PAKET BERHASIL DIUBAH")
	fmt.Println("========================================")
	fmt.Printf("No Resi        : %s\n", paket.NoResi)
	fmt.Printf("Tipe Paket     : %s\n", paket.Tipe)
	fmt.Printf("Berat          : %.2f kg\n", paket.Berat)
	fmt.Printf("Harga          : Rp %.0f\n", paket.Harga)
	fmt.Printf("Kota Pengirim  : %s\n", paket.SenderCity)
	fmt.Printf("Kota Tujuan    : %s\n", paket.ReceiverCity)
	fmt.Printf("Status         : %v\n", paket.Status)
	fmt.Printf("Dibuat Pada    : %s\n", paket.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Diperbarui Pada: %s\n", paket.UpdatedAt.Format("2006-01-02 15:04:05"))

	utils.EnterToContinue()
}
