package datas

import (
	"PackageDelivery/types"
	"fmt"
	"time"
)

var PaketDB = []types.Paket{}

func AssignPaketToKurir(username string, noResi string) bool {
	// Cari kurir berdasarkan username
	kurir := FindUserByUsername(username)
	if kurir.Username == "" {
		return false
	}

	if kurir.Role != "kurir" {
		fmt.Println("Username yang dimasukkan bukan kurir.")
		return false
	}

	// Cari paket berdasarkan noResi
	paket := CariPaket(noResi)
	if paket.NoResi == "" {
		return false
	}

	// Assign paket ke kurir
	paket.Kurir = username

	// modify usersDB
	for i, user := range UsersDB {
		if user.Username == username {
			UsersDB[i].Pakets = append(UsersDB[i].Pakets, paket)
			break
		}
	}

	// Simpan perubahan
	UpdatePaket(paket)

	return true
}

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

func UpdatePaket(paket types.Paket) {
	for i, p := range PaketDB {
		if p.NoResi == paket.NoResi {
			PaketDB[i] = paket
			return
		}
	}
}

func AddPaket(paket types.Paket) {
	for _, p := range PaketDB {
		if p.NoResi == paket.NoResi {
			fmt.Println("Paket dengan nomor resi tersebut sudah ada.")
			return
		}
	}

	now := time.Now()
	paket.CreatedAt = now
	paket.UpdatedAt = now
	paket.Status = append(paket.Status, "Paket dibuat")

	PaketDB = append(PaketDB, paket)
	fmt.Println("Paket berhasil ditambahkan.")
}
