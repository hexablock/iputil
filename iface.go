package iputil

import "net"

// GetActiveIfaces returns all interfaces in the up state
func GetActiveIfaces() ([]net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ifs := make([]net.Interface, 0, len(ifaces))
	for _, iface := range ifaces {
		// Skip interfaces that are down
		if iface.Flags&net.FlagUp != 0 {
			ifs = append(ifs, iface)
		}
	}

	return ifs, nil
}
