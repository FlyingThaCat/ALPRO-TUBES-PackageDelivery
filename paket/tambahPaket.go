package paket

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
	"strings"
	"time"
)

func InputPackageType() (string, error) {
	fmt.Println("========================================")
	fmt.Println("📦  PILIH TIPE PAKET")
	fmt.Println("========================================")
	for i, tipe := range types.PackageTypes {
		fmt.Printf("%d️⃣  %s\n", i+1, tipe)
	}
	fmt.Println("========================================")

	choice := utils.GetInt(fmt.Sprintf("Masukkan nomor tipe paket (1-%d): ", len(types.PackageTypes)), "")
	if choice < 1 || choice > len(types.PackageTypes) {
		return "", fmt.Errorf("pilihan tipe paket tidak valid")
	}

	return types.PackageTypes[choice-1], nil
}

func InputCity(prompt string) (types.Tempat, error) {
	fmt.Println("========================================")
	fmt.Printf("🏙️  PILIH %s\n", strings.ToUpper(prompt))
	fmt.Println("========================================")
	for i, kota := range types.Kota {
		fmt.Printf("%d️⃣  %s\n", i+1, kota)
	}
	fmt.Println("========================================")

	choice := utils.GetInt(fmt.Sprintf("Masukkan nomor %s (1-%d): ", strings.ToLower(prompt), len(types.Kota)), "")
	if choice < 1 || choice > len(types.Kota) {
		return types.Tempat{}, fmt.Errorf("pilihan kota tidak valid")
	}

	cityName := types.Kota[choice-1]

	for _, tempat := range types.DaftarTempat {
		if tempat.City == cityName {
			return tempat, nil
		}
	}

	return types.Tempat{}, fmt.Errorf("koordinat untuk kota %s tidak ditemukan", cityName)
}

func TambahPaket() {
	utils.ClearScreen()

	tipe, err := InputPackageType()
	if err != nil {
		fmt.Println(err.Error())
		utils.EnterToContinue()
		return
	}

	berat := utils.GetFloat("Masukkan berat paket (kg): ", "Berat tidak boleh kosong dan harus berupa angka.")

	senderCity, err := InputCity("Kota Pengirim")
	if err != nil {
		fmt.Println(err.Error())
		utils.EnterToContinue()
		return
	}

	receiverCity, err := InputCity("Kota Tujuan")
	if err != nil {
		fmt.Println(err.Error())
		utils.EnterToContinue()
		return
	}

	distance, price := utils.CalculatePackagePrice(
		senderCity.Latitude, senderCity.Longitude,
		receiverCity.Latitude, receiverCity.Longitude,
		berat, tipe,
	)

	fmt.Printf("\nJarak antar kota: %.2f km\n", distance)
	fmt.Printf("Harga paket: Rp %.0f\n", price)

	if !utils.GetConfirmation("Apakah Anda yakin ingin menambahkan paket ini?", "Silakan masukkan 'y' untuk ya atau 'n' untuk tidak.") {
		fmt.Println("Paket tidak ditambahkan.")
		utils.EnterToContinue()
		return
	}

	username := utils.GetLoggedInUsername()

	var input types.Paket
	input.Tipe = tipe
	input.Berat = berat
	input.SenderCity = types.Cities(senderCity.City)
	input.ReceiverCity = types.Cities(receiverCity.City)
	input.Harga = price
	input.NoResi = fmt.Sprintf("%d", len(datas.PaketDB)+1)
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	input.Status = []string{"Paket Dibuat"}
	input.CreatedBy = username

	datas.PaketDB = append(datas.PaketDB, input)

	fmt.Println("Paket berhasil ditambahkan.")
	DetailPaket(input)
	utils.EnterToContinue()
}
