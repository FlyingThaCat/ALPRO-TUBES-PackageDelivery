package utils

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
)

func FindPaketByUsername(username string) []types.Paket {
	var result []types.Paket
	for _, paket := range datas.PaketDB {
		if paket.Kurir == username {
			result = append(result, paket)
		}
	}
	return result
}
