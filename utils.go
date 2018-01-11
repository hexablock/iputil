package iputil

import (
	"fmt"
	"net"
	"strconv"
)

// given a advertise and bind address determine the advertise addr or return
// an error
func BuildAdvertiseAddr(a, b string) (string, string, error) {
	var addr string
	if a != "" {
		addr = a
	} else {
		// Used bind if adv is not supplied
		addr = b
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", "", err
	}

	var ipaddr *net.IPAddr
	ipaddr, err = net.ResolveIPAddr("ip", host)
	if err == nil {
		ip := ipaddr.String()

		switch ip {
		case "", "0.0.0.0", "0:0:0:0:0:0:0:0", ":", "::":
			goto INVALID_ADDR
		}

		if port == "" {
			goto INVALID_ADDR
		}

		return ip, port, nil
	}

INVALID_ADDR:
	return "", "", fmt.Errorf("invalid advertise address: %s", addr)
}

func SplitHostPort(host string) (string, int, error) {
	host, port, err := net.SplitHostPort(host)
	if err != nil {
		return "", 0, err
	}
	i, err := strconv.ParseInt(port, 10, 32)
	return host, int(i), err
}
