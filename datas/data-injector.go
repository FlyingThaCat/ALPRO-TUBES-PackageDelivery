package datas

import (
	"PackageDelivery/types"
	"time"
)

func Init() {
	now := time.Now()

	UsersDB = append(UsersDB, types.User{Username: "admin", Password: "1", Role: "admin"})
	UsersDB = append(UsersDB, types.User{
		Nama:     "kurir",
		Username: "kurir",
		Password: "1",
		Role:     "kurir",
		Pakets: []types.Paket{
			{
				NoResi:       "5",
				Tipe:         "SameDay",
				Berat:        5,
				Harga:        150000,
				SenderCity:   types.Jakarta,
				ReceiverCity: types.Bekasi,
				Status:       []string{"Paket Dibuat"},
				CreatedAt:    now,
				UpdatedAt:    now,
				Kurir:        "kurir",
			},
		},
	})
	UsersDB = append(UsersDB, types.User{Username: "user", Password: "1", Role: "user", Pakets: []types.Paket{}})

	PaketDB = append(PaketDB,
		types.Paket{
			NoResi:       "1",
			Tipe:         "Reguler",
			Berat:        1,
			Harga:        10000,
			SenderCity:   types.Jakarta,
			ReceiverCity: types.Depok,
			Status:       []string{"Paket Dibuat"},
			CreatedAt:    now,
			UpdatedAt:    now,
			Kurir:        "",
		},
		types.Paket{
			NoResi:       "2",
			Tipe:         "Express",
			Berat:        5,
			Harga:        80000,
			SenderCity:   types.Bekasi,
			ReceiverCity: types.Bogor,
			Status:       []string{"Paket Dibuat"},
			CreatedAt:    now,
			UpdatedAt:    now,
			Kurir:        "",
		},
		types.Paket{
			NoResi:       "3",
			Tipe:         "SameDay",
			Berat:        2,
			Harga:        60000,
			SenderCity:   types.Depok,
			ReceiverCity: types.Bogor,
			Status:       []string{"Paket Dibuat"},
			CreatedAt:    now,
			UpdatedAt:    now,
			Kurir:        "",
		},
		types.Paket{
			NoResi:       "4",
			Tipe:         "Reguler",
			Berat:        4,
			Harga:        50000,
			SenderCity:   types.Jakarta,
			ReceiverCity: types.Bekasi,
			Status:       []string{"Paket Dibuat"},
			CreatedAt:    now,
			UpdatedAt:    now,
			Kurir:        "",
		},
		types.Paket{
			NoResi:       "5",
			Tipe:         "SameDay",
			Berat:        5,
			Harga:        150000,
			SenderCity:   types.Jakarta,
			ReceiverCity: types.Bekasi,
			Status:       []string{"Paket Dibuat"},
			CreatedAt:    now,
			UpdatedAt:    now,
			Kurir:        "kurir",
		},
	)
}
