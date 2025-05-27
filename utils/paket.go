package utils

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"fmt"
	"time"
)

func FindPaketByNoResi(noResi string) *types.Paket {
	for i := range datas.PaketDB {
		if datas.PaketDB[i].NoResi == noResi {
			return &datas.PaketDB[i]
		}
	}
	return nil
}

func FindPaketByUsername(username string) []types.Paket {
	var result []types.Paket
	for _, paket := range datas.PaketDB {
		if paket.Kurir == username {
			result = append(result, paket)
		}
	}
	return result
}

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
	paket := FindPaketByNoResi(noResi)
	if paket.NoResi == "" {
		return false
	}

	// Assign paket ke kurir
	paket.Kurir = username

	// modify usersDB
	for i, user := range datas.UsersDB {
		if user.Username == username {
			datas.UsersDB[i].Pakets = append(datas.UsersDB[i].Pakets, *paket)
			break
		}
	}

	// Simpan perubahan
	UpdatePaket(*paket)

	return true
}

func DeletePaket(noResi string) bool {
	// Hapus paket dari database
	for i, p := range datas.PaketDB {
		if p.NoResi == noResi {
			datas.PaketDB = append(datas.PaketDB[:i], datas.PaketDB[i+1:]...)
			return true
		}
	}
	return false
}

func UpdatePaket(paket types.Paket) {
	for i, p := range datas.PaketDB {
		if p.NoResi == paket.NoResi {
			datas.PaketDB[i] = paket
			return
		}
	}
}

func AddPaket(paket types.Paket) {
	for _, p := range datas.PaketDB {
		if p.NoResi == paket.NoResi {
			fmt.Println("Paket dengan nomor resi tersebut sudah ada.")
			return
		}
	}

	now := time.Now()
	paket.CreatedAt = now
	paket.UpdatedAt = now
	paket.Status = append(paket.Status, "Paket dibuat")

	datas.PaketDB = append(datas.PaketDB, paket)
	fmt.Println("Paket berhasil ditambahkan.")
}