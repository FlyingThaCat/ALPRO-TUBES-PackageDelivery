package utils

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
)

func FindUserByUsername(username string) types.User {
	for _, user := range datas.UsersDB {
		if user.Username == username {
			return user
		}
	}
	return types.User{}
}