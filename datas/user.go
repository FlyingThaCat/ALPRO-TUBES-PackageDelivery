package datas

import (
	"PackageDelivery/types"
)

var UsersDB = []types.User{}

func FindUserByUsername(username string) types.User {
	for _, user := range UsersDB {
		if user.Username == username {
			return user
		}
	}
	return types.User{}
}