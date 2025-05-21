package utils

var loggedInUsername string

func SetLoggedInUsername(username string) {
	loggedInUsername = username
}

func GetLoggedInUsername() string {
	return loggedInUsername
}
