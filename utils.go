package iputil

import (
	"fmt"
	"net"
	"strconv"
)

// BuildAdvertiseAddr returns the advertise host and port string based on the given
// advertise and bind addresses or an error
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

// SplitHostPort splits a host port string and returns a host string
// and integer port or an error
func SplitHostPort(host string) (string, int, error) {
	host, port, err := net.SplitHostPort(host)
	if err != nil {
		return "", 0, err
	}
	i, err := strconv.ParseInt(port, 10, 32)
	return host, int(i), err
}

// PortFromString returns the port number or an error from the given string
func PortFromString(port string) (int, error) {
	i, err := strconv.ParseInt(port, 10, 32)
	return int(i), err
}
