package utils

import (
	"PackageDelivery/types"
	"fmt"
)

const PricePerKm = 5000.0

var PaketTypeMultiplier = map[string]float64{
	"reguler": 1.0,
	"express": 1.25,
	"sameday": 1.5,
}

func GetCoordinates(city types.Cities) (lat, lon float64, err error) {
	for _, tempat := range types.DaftarTempat {
		if tempat.City == string(city) {
			return tempat.Latitude, tempat.Longitude, nil
		}
	}
	return 0, 0, fmt.Errorf("koordinat untuk kota %s tidak ditemukan", city)
}

func CalculatePackagePrice(senderLat, senderLon, receiverLat, receiverLon float64, berat float64, tipe string) (float64, float64) {
	distance := HitungTwoPoints(senderLat, senderLon, receiverLat, receiverLon)
	basePrice := distance * PricePerKm

	if berat > 5 {
		basePrice += (berat - 5) * 1000
	}

	multiplier, ok := PaketTypeMultiplier[tipe]
	if !ok {
		multiplier = 1.0
	}
	totalPrice := basePrice * multiplier

	return distance, totalPrice
}
