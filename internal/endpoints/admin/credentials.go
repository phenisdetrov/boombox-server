package admin

import (
	"github.com/phenisdetrov/boombox-server/internal/config"
)

func credentials(username, password string) bool {
	if !config.UsernameRegexp.MatchString(username) || !config.PasswordRegexp.MatchString(password) {
		return false
	}
	return true
}
