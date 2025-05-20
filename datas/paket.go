package datas

import "PackageDelivery/types"

var PaketDB = []types.Paket{}

func CariPaket(noResi string) types.Paket {
	for _, paket := range PaketDB {
		if paket.NoResi == noResi {
			return paket
		}
	}
	return types.Paket{}
}

func HapusPaket(noResi string) bool {
	// Hapus paket dari database
	for i, p := range PaketDB {
		if p.NoResi == noResi {
			PaketDB = append(PaketDB[:i], PaketDB[i+1:]...)
			return true
		}
	}
	return false
}