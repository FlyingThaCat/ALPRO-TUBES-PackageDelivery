package paket

import "PackageDelivery/datas"

func HapusPaket(noResi string) bool {
	// Hapus paket dari database
	for i, p := range datas.PaketDB {
		if p.NoResi == noResi {
			datas.PaketDB = append(datas.PaketDB[:i], datas.PaketDB[i+1:]...)
			return true
		}
	}
	return false
}