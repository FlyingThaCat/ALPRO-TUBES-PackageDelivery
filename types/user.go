package types

type User struct {
	Nama     string
	Username string
	Password string
	Role     string // "user", "kurir", "admin"
	Pakets   []Paket
}