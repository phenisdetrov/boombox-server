package get

import (
    "regexp"
)

const LISTEN = "10.0.0.1:20000"
const VERSION = "0.0.1"

var UsernameRegexp = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)
var PasswordRegexp = regexp.MustCompile(`^[\x20-\x7E]+$`)
