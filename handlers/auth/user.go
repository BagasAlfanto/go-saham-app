package auth

import (
	"fmt"
	"saham-app/model/user"
)

func Authentic(username, password string) bool {
	success, message := user.CheckPassword(username, password)

	fmt.Println(message)
	return success
}

