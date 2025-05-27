package utils

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
)

func AddUser(user types.User) {
	datas.UsersDB = append(datas.UsersDB, user)
}

func DeleteUser(username string) bool {
	for i, p := range datas.UsersDB {
		if p.Username == username {
			datas.UsersDB = append(datas.UsersDB[:i], datas.UsersDB[i+1:]...)
			return true
		}
	}
	return false
}

func FindUserByUsername(username string) types.User {
	for _, user := range datas.UsersDB {
		if user.Username == username {
			return user
		}
	}
	return types.User{}
}