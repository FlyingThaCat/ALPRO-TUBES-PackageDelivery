package paket

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
)

func CariPaket(noResi string) types.Paket {
	for _, paket := range datas.PaketDB {
		if paket.NoResi == noResi {
			return paket
		}
	}
	return types.Paket{}
}