package config

import (
	"regexp"
)

const LISTEN = "10.0.0.1:20000"
const VERSION = "0.0.1"
const POSTGRES = "postgres://admin:admin@127.0.0.1:5432/boombox"

var UsernameRegexp = regexp.MustCompile(`^[A-Za-z0-9._-]{1,72}$`)
var PasswordRegexp = regexp.MustCompile(`^[\x20-\x7E]{1,72}$`)
