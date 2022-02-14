package formatter

import (
	"os/user"
)

// Current is used only in testing and mocking.
var Current = user.Current // nolint: gochecknoglobals

func getUser() string {
	if u, err := Current(); err == nil {
		return u.Username
	}

	return ""
}
