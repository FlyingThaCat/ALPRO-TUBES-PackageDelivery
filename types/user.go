package types

type User struct {
	Username string
	Password string
	Role     string // "user", "kurir", "admin"
	Pakets   []Paket
}