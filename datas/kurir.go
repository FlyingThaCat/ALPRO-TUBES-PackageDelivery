package datas

func HapusKurir(username string) bool {
	// Hapus Kurir dari database
	for i, p := range UsersDB {
		if p.Username == username {
			UsersDB = append(UsersDB[:i], UsersDB[i+1:]...)
			return true
		}
	}
	return false
}