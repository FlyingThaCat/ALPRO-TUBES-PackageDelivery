package datas

import (
	"PackageDelivery/types"
	"time"
)

func Init() {
	UsersDB = append(UsersDB, types.User{Username: "admin", Password: "1", Role: "admin"})
	UsersDB = append(UsersDB, types.User{Username: "kurir", Password: "1", Role: "kurir"})
	UsersDB = append(UsersDB, types.User{Username: "user", Password: "1", Role: "user"})

	PaketDB = append(PaketDB, types.Paket{NoResi: "1", Tipe: "Reguler", Berat: 1, Harga: 10000, SenderCity: "Jakarta", ReceiverCity: "Depok", Status: []string{"Paket Dibuat"}, CreatedAt: time.Now(), UpdatedAt: time.Now(), Kurir: "" })
	PaketDB = append(PaketDB, types.Paket{NoResi: "2", Tipe: "Express", Berat: 5, Harga: 80000, SenderCity: "Bogor", ReceiverCity: "Tangerang", Status: []string{"Paket Dibuat"}, CreatedAt: time.Now(), UpdatedAt: time.Now(), Kurir: "" })
	PaketDB = append(PaketDB, types.Paket{NoResi: "3", Tipe: "SameDay", Berat: 2, Harga: 60000, SenderCity: "Depok", ReceiverCity: "Bogor", Status: []string{"Paket Dibuat"}, CreatedAt: time.Now(), UpdatedAt: time.Now(), Kurir: "" })
	PaketDB = append(PaketDB, types.Paket{NoResi: "4", Tipe: "Reguler", Berat: 4, Harga: 50000, SenderCity: "Jakarta", ReceiverCity: "Bekasi", Status: []string{"Paket Dibuat"}, CreatedAt: time.Now(), UpdatedAt: time.Now(), Kurir: "" })
	PaketDB = append(PaketDB, types.Paket{NoResi: "5", Tipe: "SameDay", Berat: 5, Harga: 150000, SenderCity: "Jakarta", ReceiverCity: "Bekasi", Status: []string{"Paket Dibuat"}, CreatedAt: time.Now(), UpdatedAt: time.Now(), Kurir: "" })
}