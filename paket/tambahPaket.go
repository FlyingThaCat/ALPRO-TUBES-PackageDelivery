package paket

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"fmt"
	"strings"
	"time"
)

const PricePerKm = 5000.0

var PaketTypeMultiplier = map[string]float64{
	"reguler": 1.0,
	"express": 1.25,
	"sameday": 1.5,
}

func CalculatePackagePrice(senderLat, senderLon, receiverLat, receiverLon float64, berat float64, tipe string) (float64, float64) {
	distance := utils.HitungTwoPoints(senderLat, senderLon, receiverLat, receiverLon)
	basePrice := distance * PricePerKm

	if berat > 5 {
		basePrice += (berat - 5) * 1000
	}

	multiplier, ok := PaketTypeMultiplier[strings.ToLower(tipe)]
	if !ok {
		multiplier = 1.0
	}
	totalPrice := basePrice * multiplier

	return distance, totalPrice
}

// InputPackageType meminta input dari user terkait tipe paket
func InputPackageType() (string, error) {
	fmt.Println("Tipe Paket:")
	for _, tipe := range types.PackageTypes {
		fmt.Printf("- %s\n", tipe)
	}
	tipe := utils.GetInput("Masukkan tipe paket: ", false, "Tipe paket tidak boleh kosong.")

	if !types.IsValidPackageType(tipe) {
		return "", fmt.Errorf("tipe paket tidak valid")
	}
	return tipe, nil
}

// InputCity meminta input nama kota dan validasi keberadaannya
func InputCity(prompt string) (types.Tempat, error) {
	fmt.Println(prompt)
	for _, kota := range types.Kota {
		fmt.Printf("- %s\n", kota)
	}
	fmt.Printf("Masukkan %s: ", strings.ToLower(prompt))

	cityName := utils.GetInput("", false, "Kota tidak boleh kosong.")

	valid := false
	for _, kota := range types.Kota {
		if strings.EqualFold(kota, cityName) {
			valid = true
			cityName = kota
			break
		}
	}
	if !valid {
		return types.Tempat{}, fmt.Errorf("kota %s tidak ditemukan", cityName)
	}

	for _, tempat := range types.DaftarTempat {
		if tempat.City == cityName {
			return tempat, nil
		}
	}

	return types.Tempat{}, fmt.Errorf("koordinat untuk kota %s tidak ditemukan", cityName)
}

// TambahPaket menambahkan data paket baru dengan harga berdasarkan jarak, berat, dan tipe
func TambahPaket() {
	utils.ClearScreen()

	tipe, err := InputPackageType()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	berat := utils.GetFloat("Masukkan berat paket (kg): ", "Berat tidak boleh kosong dan harus berupa angka.")

	senderCity, err := InputCity("Kota Pengirim")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	receiverCity, err := InputCity("Kota Tujuan")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	distance, price := CalculatePackagePrice(
		senderCity.Latitude, senderCity.Longitude,
		receiverCity.Latitude, receiverCity.Longitude,
		berat, tipe,
	)

	fmt.Printf("\nJarak antar kota: %.2f km\n", distance)
	fmt.Printf("Harga paket: Rp %.0f\n", price)

	if !utils.GetConfirmation("Apakah Anda yakin ingin menambahkan paket ini?", "Silakan masukkan 'y' untuk ya atau 'n' untuk tidak.") {
		fmt.Println("Paket tidak ditambahkan.")
		return
	}

	// Buat paket baru
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

	datas.PaketDB = append(datas.PaketDB, input)

	fmt.Println("Paket berhasil ditambahkan.")
	DetailPaket(input)
}
