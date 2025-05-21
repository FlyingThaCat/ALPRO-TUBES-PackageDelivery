package utils

import "math"

// Haversine menghitung jarak antara dua titik koordinat (latitude, longitude) dalam kilometer
func HitungTwoPoints(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius bumi dalam kilometer

	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// CalculatePrice menghitung harga berdasarkan jarak dan tarif per kilometer
func CalculatePrice(lat1, lon1, lat2, lon2, pricePerKm float64) float64 {
	distance := HitungTwoPoints(lat1, lon1, lat2, lon2)
	price := distance * pricePerKm
	return price
}
