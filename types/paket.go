package types

import "time"

type Paket struct {
	NoResi       string
	Tipe         string
	Berat        float64
	Harga        float64
	SenderCity   Cities
	ReceiverCity Cities
	Status       []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Kurir        string
	CreatedBy    string
	ReceiverName string
}
