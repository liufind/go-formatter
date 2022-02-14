package formatter

import (
	"net"
)

// Dial is used only in testing and mocking.
var Dial = net.Dial // nolint: gochecknoglobals

func getIPAddress() string {
	connection, err := Dial("udp", "8.8.8.8:80")

	if connection == nil {
		return "127.0.0.1"
	}

	defer func() {
		err = connection.Close()
		_ = err
	}()

	return connection.LocalAddr().(*net.UDPAddr).IP.String()
}
