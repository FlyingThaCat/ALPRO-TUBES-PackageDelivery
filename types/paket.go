package types

import "time"

type Paket struct {
	NoResi       string
	Tipe         string
	Berat        float64
	Harga        float64
	SenderCity   string
	ReceiverCity string
	Status       []string // Paket Dibuat, Paket Telah Diserahkan ke kurir, Paket dalam Perjalanan, Paket Telah Sampai di tujuan
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Kurir        string
}
