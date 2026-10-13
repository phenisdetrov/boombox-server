package admin

import (
	"github.com/phenisdetrov/boombox-server/internal/config"
)

func checkCredentials(username, password string) bool {
	if !config.UsernameRegexp.MatchString(username) || !config.PasswordRegexp.MatchString(password) {
		return false
	}
	if len(username) > 72 || len(password) > 72 {
		return false
	}
	return true
}
