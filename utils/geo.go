package utils

import (
	"PackageDelivery/types"
	"math"
)

func GetJarak(from, to types.Cities) float64 {
	var lat1, lon1, lat2, lon2 float64

	for _, tempat := range types.DaftarTempat {
		if tempat.City == string(from) {
			lat1 = tempat.Latitude
			lon1 = tempat.Longitude
		}
		if tempat.City == string(to) {
			lat2 = tempat.Latitude
			lon2 = tempat.Longitude
		}
	}

	return Haversine(lat1, lon1, lat2, lon2)
}

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius bumi dalam km
	dLat := toRadians(lat2 - lat1)
	dLon := toRadians(lon2 - lon1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(toRadians(lat1))*math.Cos(toRadians(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func toRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
