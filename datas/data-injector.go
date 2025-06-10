package datas

import (
	"PackageDelivery/types"
	"time"
)

// Init populates UsersDB and PaketDB with realistic, relational test data
func Init() {
	now := time.Now()

	// Admin user
	UsersDB = append(UsersDB, types.User{
		Nama:     "Administrator",
		Username: "admin",
		Password: "admin123",
		Role:     "admin",
	})

	// Courier users with their own Pakets
	UsersDB = append(UsersDB, types.User{
		Nama:     "Budi Santoso",
		Username: "budi",
		Password: "kurirBudi",
		Role:     "kurir",
		Pakets: []types.Paket{
			{NoResi: "1", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Jakarta, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-1 * time.Hour), UpdatedAt: now, Kurir: "budi", CreatedBy: "admin", ReceiverName: "komar"},	
			{NoResi: "6", Tipe: "SameDay", Berat: 3.5, Harga: 65000, SenderCity: types.Depok, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-6 * time.Hour), UpdatedAt: now.Add(-2 * time.Hour), Kurir: "budi", CreatedBy: "admin", ReceiverName: "komar"},	
			{NoResi: "8", Tipe: "Express", Berat: 1.5, Harga: 25000, SenderCity: types.Tangerang, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-8 * time.Hour), UpdatedAt: now.Add(-4 * time.Hour), Kurir: "budi", CreatedBy: "rina", ReceiverName: "komar"},	
			{NoResi: "11", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Depok, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-11 * time.Hour), UpdatedAt: now.Add(-6 * time.Hour), Kurir: "budi", CreatedBy: "rina", ReceiverName: "komar"},	
			{NoResi: "16", Tipe: "Reguler", Berat: 2.5, Harga: 25000, SenderCity: types.Bogor, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-16 * time.Hour), UpdatedAt: now.Add(-9 * time.Hour), Kurir: "budi", CreatedBy: "rina", ReceiverName: "komar"},	
			{NoResi: "22", Tipe: "Reguler", Berat: 2.5, Harga: 25000, SenderCity: types.Depok, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-22 * time.Hour), UpdatedAt: now.Add(-13 * time.Hour), Kurir: "budi", CreatedBy: "dedi", ReceiverName: "komar"},	
		},
	})

	UsersDB = append(UsersDB, types.User{
		Nama:     "Siti Aminah",
		Username: "siti",
		Password: "kurirSiti",
		Role:     "kurir",
		Pakets: []types.Paket{
			{NoResi: "3", Tipe: "SameDay", Berat: 2.0, Harga: 50000, SenderCity: types.Bekasi, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil", "Dalam Pengiriman"}, CreatedAt: now.Add(-3 * time.Hour), UpdatedAt: now.Add(-1 * time.Hour), Kurir: "siti", CreatedBy: "admin", ReceiverName: "tara"},
			{NoResi: "13", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Tangerang, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-13 * time.Hour), UpdatedAt: now.Add(-7 * time.Hour), Kurir: "siti", CreatedBy: "dedi", ReceiverName: "tara"},
			{NoResi: "19", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Tangerang, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-19 * time.Hour), UpdatedAt: now.Add(-11 * time.Hour), Kurir: "siti", CreatedBy: "dedi", ReceiverName: "tara"},
			{NoResi: "23", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Bogor, ReceiverCity: types.Tangerang, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-23 * time.Hour), UpdatedAt: now.Add(-14 * time.Hour), Kurir: "siti", CreatedBy: "dedi", ReceiverName: "tara"},
		},
	})
	UsersDB = append(UsersDB, types.User{
		Nama:     "Anton Wijaya",
		Username: "anton",
		Password: "kurirAnton",
		Role:     "kurir",
		Pakets: []types.Paket{
			{NoResi: "5", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Jakarta, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-5 * time.Hour), UpdatedAt: now.Add(-3 * time.Hour), Kurir: "anton", CreatedBy: "budi", ReceiverName: "putra"},
			{NoResi: "7", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Bogor, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil", "Dalam Pengiriman"}, CreatedAt: now.Add(-7 * time.Hour), UpdatedAt: now.Add(-1 * time.Hour), Kurir: "anton", CreatedBy: "budi", ReceiverName: "putra"},
			{NoResi: "14", Tipe: "Express", Berat: 1.5, Harga: 25000, SenderCity: types.Jakarta, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-14 * time.Hour), UpdatedAt: now.Add(-8 * time.Hour), Kurir: "anton", CreatedBy: "budi", ReceiverName: "putra"},
			{NoResi: "17", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Bekasi, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-17 * time.Hour), UpdatedAt: now.Add(-10 * time.Hour), Kurir: "anton", CreatedBy: "budi", ReceiverName: "tara"},
			{NoResi: "20", Tipe: "Express", Berat: 1.5, Harga: 25000, SenderCity: types.Jakarta, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-20 * time.Hour), UpdatedAt: now.Add(-12 * time.Hour), Kurir: "anton", CreatedBy: "budi", ReceiverName: "tara"},
			{NoResi: "25", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Tangerang, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-25 * time.Hour), UpdatedAt: now.Add(-15 * time.Hour), Kurir: "anton", CreatedBy: "budi", ReceiverName: "tara"},
		},
	})

	UsersDB = append(UsersDB, types.User{
		Nama:     "Lina Putri",
		Username: "lina",
		Password: "kurirLina",
		Role:     "kurir",
		Pakets: []types.Paket{
			{NoResi: "10", Tipe: "Reguler", Berat: 2.5, Harga: 25000, SenderCity: types.Bekasi, ReceiverCity: types.Tangerang, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-10 * time.Hour), UpdatedAt: now.Add(-5 * time.Hour), Kurir: "lina", CreatedBy: "dedi", ReceiverName: "budiman"},
			{NoResi: "15", Tipe: "SameDay", Berat: 2.0, Harga: 50000, SenderCity: types.Depok, ReceiverCity: types.Tangerang, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-15 * time.Hour), UpdatedAt: now.Add(-15 * time.Hour), Kurir: "lina", CreatedBy: "admin", ReceiverName: "budiman"},
		},
	})

	// Regular users
	UsersDB = append(UsersDB, types.User{Nama: "Rina Wahyuni", Username: "rina", Password: "userRina", Role: "user", Pakets: []types.Paket{}})
	UsersDB = append(UsersDB, types.User{Nama: "Dedi Pratama", Username: "dedi", Password: "userDedi", Role: "user", Pakets: []types.Paket{}})

	// All PaketDB entries (including unassigned)
	PaketDB = []types.Paket{
		{NoResi: "1", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Jakarta, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-1 * time.Hour), UpdatedAt: now, Kurir: "budi", CreatedBy: "admin", ReceiverName: "siti"},
		{NoResi: "2", Tipe: "Express", Berat: 1.5, Harga: 25000, SenderCity: types.Depok, ReceiverCity: types.Tangerang, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-2 * time.Hour), UpdatedAt: now.Add(-2 * time.Hour), Kurir: "", CreatedBy: "admin", ReceiverName: "siti"},
		{NoResi: "3", Tipe: "SameDay", Berat: 2.0, Harga: 50000, SenderCity: types.Bekasi, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil", "Dalam Pengiriman"}, CreatedAt: now.Add(-3 * time.Hour), UpdatedAt: now.Add(-1 * time.Hour), Kurir: "siti", CreatedBy: "admin", ReceiverName: "danang"},
		{NoResi: "4", Tipe: "Reguler", Berat: 2.5, Harga: 25000, SenderCity: types.Tangerang, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-4 * time.Hour), UpdatedAt: now.Add(-4 * time.Hour), Kurir: "", CreatedBy: "admin", ReceiverName: "danang"},
		{NoResi: "5", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Jakarta, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-5 * time.Hour), UpdatedAt: now.Add(-3 * time.Hour), Kurir: "anton", CreatedBy: "admin", ReceiverName: "danang"},
		{NoResi: "6", Tipe: "SameDay", Berat: 3.5, Harga: 65000, SenderCity: types.Depok, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-6 * time.Hour), UpdatedAt: now.Add(-2 * time.Hour), Kurir: "budi", CreatedBy: "admin", ReceiverName: "danang"},
		{NoResi: "7", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Bogor, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil", "Dalam Pengiriman"}, CreatedAt: now.Add(-7 * time.Hour), UpdatedAt: now.Add(-1 * time.Hour), Kurir: "anton", CreatedBy: "admin", ReceiverName: "budiman"},
		{NoResi: "8", Tipe: "Express", Berat: 1.5, Harga: 25000, SenderCity: types.Tangerang, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-8 * time.Hour), UpdatedAt: now.Add(-4 * time.Hour), Kurir: "budi", CreatedBy: "admin", ReceiverName: "budiman"},
		{NoResi: "9", Tipe: "SameDay", Berat: 2.0, Harga: 50000, SenderCity: types.Jakarta, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-9 * time.Hour), UpdatedAt: now.Add(-9 * time.Hour), Kurir: "", CreatedBy: "admin", ReceiverName: "budiman"},
		{NoResi: "10", Tipe: "Reguler", Berat: 2.5, Harga: 25000, SenderCity: types.Bekasi, ReceiverCity: types.Tangerang, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-10 * time.Hour), UpdatedAt: now.Add(-5 * time.Hour), Kurir: "lina", CreatedBy: "admin", ReceiverName: "budiman"},
		{NoResi: "11", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Depok, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-11 * time.Hour), UpdatedAt: now.Add(-6 * time.Hour), Kurir: "budi", CreatedBy: "rina", ReceiverName: "siti"},
		{NoResi: "12", Tipe: "SameDay", Berat: 3.5, Harga: 65000, SenderCity: types.Bogor, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-12 * time.Hour), UpdatedAt: now.Add(-12 * time.Hour), Kurir: "", CreatedBy: "rina", ReceiverName: "siti"},
		{NoResi: "13", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Tangerang, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-13 * time.Hour), UpdatedAt: now.Add(-7 * time.Hour), Kurir: "", CreatedBy: "rina", ReceiverName: "komar"},
		{NoResi: "14", Tipe: "Express", Berat: 1.5, Harga: 25000, SenderCity: types.Jakarta, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-14 * time.Hour), UpdatedAt: now.Add(-8 * time.Hour), Kurir: "anton", CreatedBy: "rina", ReceiverName: "komar"},
		{NoResi: "15", Tipe: "SameDay", Berat: 2.0, Harga: 50000, SenderCity: types.Depok, ReceiverCity: types.Tangerang, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-15 * time.Hour), UpdatedAt: now.Add(-15 * time.Hour), Kurir: "lina", CreatedBy: "rina", ReceiverName: "komar"},
		{NoResi: "16", Tipe: "Reguler", Berat: 2.5, Harga: 25000, SenderCity: types.Bogor, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-16 * time.Hour), UpdatedAt: now.Add(-9 * time.Hour), Kurir: "budi", CreatedBy: "rina", ReceiverName: "siti"},
		{NoResi: "17", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Bekasi, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-17 * time.Hour), UpdatedAt: now.Add(-10 * time.Hour), Kurir: "", CreatedBy: "rina", ReceiverName: "siti"},
		{NoResi: "18", Tipe: "SameDay", Berat: 3.5, Harga: 65000, SenderCity: types.Depok, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-18 * time.Hour), UpdatedAt: now.Add(-18 * time.Hour), Kurir: "", CreatedBy: "admin", ReceiverName: "siti"},
		{NoResi: "19", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Tangerang, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-19 * time.Hour), UpdatedAt: now.Add(-11 * time.Hour), Kurir: "siti", CreatedBy: "admin", ReceiverName: "siti"},
		{NoResi: "20", Tipe: "Express", Berat: 1.5, Harga: 25000, SenderCity: types.Jakarta, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-20 * time.Hour), UpdatedAt: now.Add(-12 * time.Hour), Kurir: "anton", CreatedBy: "admin", ReceiverName: "siti"},
		{NoResi: "21", Tipe: "SameDay", Berat: 2.0, Harga: 50000, SenderCity: types.Bekasi, ReceiverCity: types.Bogor, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-21 * time.Hour), UpdatedAt: now.Add(-21 * time.Hour), Kurir: "", CreatedBy: "admin", ReceiverName: "siti"},
		{NoResi: "22", Tipe: "Reguler", Berat: 2.5, Harga: 25000, SenderCity: types.Depok, ReceiverCity: types.Jakarta, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-22 * time.Hour), UpdatedAt: now.Add(-13 * time.Hour), Kurir: "budi", CreatedBy: "dedi", ReceiverName: "tara"},
		{NoResi: "23", Tipe: "Express", Berat: 3.0, Harga: 40000, SenderCity: types.Bogor, ReceiverCity: types.Tangerang, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-23 * time.Hour), UpdatedAt: now.Add(-14 * time.Hour), Kurir: "", CreatedBy: "dedi", ReceiverName: "tara"},
		{NoResi: "24", Tipe: "SameDay", Berat: 3.5, Harga: 65000, SenderCity: types.Jakarta, ReceiverCity: types.Bekasi, Status: []string{"Paket Dibuat"}, CreatedAt: now.Add(-24 * time.Hour), UpdatedAt: now.Add(-24 * time.Hour), Kurir: "", CreatedBy: "dedi", ReceiverName: "tara"},
		{NoResi: "25", Tipe: "Reguler", Berat: 1.0, Harga: 10000, SenderCity: types.Tangerang, ReceiverCity: types.Depok, Status: []string{"Paket Dibuat", "Diambil"}, CreatedAt: now.Add(-25 * time.Hour), UpdatedAt: now.Add(-15 * time.Hour), Kurir: "anton", CreatedBy: "dedi", ReceiverName: "siti"},
	}
}
