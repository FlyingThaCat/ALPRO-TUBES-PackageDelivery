package datas

import "PackageDelivery/types"

func Init() {
	UsersDB = append(UsersDB, types.User{Username: "admin", Password: "1", Role: "admin"})
	UsersDB = append(UsersDB, types.User{Username: "kurir", Password: "1", Role: "kurir"})
	UsersDB = append(UsersDB, types.User{Username: "user", Password: "1", Role: "user"})
}