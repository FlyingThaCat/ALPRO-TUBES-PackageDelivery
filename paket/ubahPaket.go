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
	fmt.Println("✏  UBAH DATA PAKET")
	fmt.Println("========================================")

	// Input Tipe Paket (1. Reguler, 2. Express, 3. Sameday)
	fmt.Println("Tipe Paket:")
	fmt.Println("1. Reguler")
	fmt.Println("2. Express")
	fmt.Println("3. Sameday")

	currentType := strings.ToLower(string(paket.Tipe))
	currentOption := "1"
	switch currentType {
	case "reguler":
		currentOption = "1"
	case "express":
		currentOption = "2"
	case "sameday":
		currentOption = "3"
	}

	fmt.Printf("Pilih Tipe Paket [default %s]: ", currentOption)
	typeInput, _ := reader.ReadString('\n')
	typeInput = strings.TrimSpace(typeInput)
	if typeInput == "" {
		typeInput = currentOption
	}

	switch typeInput {
	case "1":
		paket.Tipe = "reguler"
	case "2":
		paket.Tipe = "express"
	case "3":
		paket.Tipe = "sameday"
	default:
		fmt.Println("\n⚠  Input tidak valid. Menggunakan nilai sebelumnya.")
	}

	// Input Berat
	fmt.Printf("Berat (kg) [%.2f]: ", paket.Berat)
	beratStr, _ := reader.ReadString('\n')
	beratStr = strings.TrimSpace(beratStr)
	if beratStr != "" {
		fmt.Sscanf(beratStr, "%f", &paket.Berat)
	}

	// Pilih Kota Pengirim
	fmt.Println("\nDaftar Kota:")
	for i, city := range types.AllCities() {
		fmt.Printf("%d. %s\n", i+1, city)
	}

	currentSenderIndex := 1
	for i, city := range types.AllCities() {
		if string(paket.SenderCity) == city {
			currentSenderIndex = i + 1
			break
		}
	}
	fmt.Printf("Pilih Kota Pengirim [default %d]: ", currentSenderIndex)
	senderInput, _ := reader.ReadString('\n')
	senderInput = strings.TrimSpace(senderInput)
	if senderInput == "" {
		senderInput = fmt.Sprintf("%d", currentSenderIndex)
	}

	senderIndex := 0
	_, err := fmt.Sscanf(senderInput, "%d", &senderIndex)
	if err != nil || senderIndex < 1 || senderIndex > len(types.AllCities()) {
		fmt.Println("\n⚠  Input tidak valid.")
		utils.EnterToContinue()
		return
	}
	paket.SenderCity = types.ParseCity(types.AllCities()[senderIndex-1])

	// Pilih Kota Tujuan
	fmt.Println("\nDaftar Kota:")
	for i, city := range types.AllCities() {
		fmt.Printf("%d. %s\n", i+1, city)
	}

	currentReceiverIndex := 1
	for i, city := range types.AllCities() {
		if string(paket.ReceiverCity) == city {
			currentReceiverIndex = i + 1
			break
		}
	}
	fmt.Printf("Pilih Kota Tujuan [default %d]: ", currentReceiverIndex)
	receiverInput, _ := reader.ReadString('\n')
	receiverInput = strings.TrimSpace(receiverInput)
	if receiverInput == "" {
		receiverInput = fmt.Sprintf("%d", currentReceiverIndex)
	}

	receiverIndex := 0
	_, err = fmt.Sscanf(receiverInput, "%d", &receiverIndex)
	if err != nil || receiverIndex < 1 || receiverIndex > len(types.AllCities()) {
		fmt.Println("\n⚠  Input tidak valid.")
		utils.EnterToContinue()
		return
	}
	paket.ReceiverCity = types.ParseCity(types.AllCities()[receiverIndex-1])

	// Recalculate Price
	senderLat, senderLon, err := utils.GetCoordinates(paket.SenderCity)
	if err != nil {
		fmt.Printf("\n⚠ Error koordinat kota pengirim: %s\n", err)
		utils.EnterToContinue()
		return
	}

	receiverLat, receiverLon, err := utils.GetCoordinates(paket.ReceiverCity)
	if err != nil {
		fmt.Printf("\n⚠ Error koordinat kota tujuan: %s\n", err)
		utils.EnterToContinue()
		return
	}

	_, newPrice := utils.CalculatePackagePrice(senderLat, senderLon, receiverLat, receiverLon, paket.Berat, strings.ToLower(paket.Tipe))
	paket.Harga = newPrice
	paket.UpdatedAt = time.Now()

	// Update Paket di Database
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
