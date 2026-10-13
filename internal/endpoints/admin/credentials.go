package admin

import (
	"github.com/phenisdetrov/boombox-server/internal/config"
)

func checkCredentials(username, password string) bool {
	return config.UsernameRegexp.MatchString(username) && config.PasswordRegexp.MatchString(password)
}
