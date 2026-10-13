package admin

import (
    "github.com/phenisdetrov/boombox-server/internal/get"
)

func CheckCredentials(username, password string) bool {
	if !get.UsernameRegexp.MatchString(username) || !get.PasswordRegexp.MatchString(password) {
		return false
	}
	if len(username) > 72 || len(password) > 72 {
		return false
	}
	return true
}
