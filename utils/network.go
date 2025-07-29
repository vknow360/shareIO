package utils

import (
	"net"
	"strings"
)

func GetLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "localhost"
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() || ip.To4() == nil {
				continue
			}

			ipStr := ip.String()

			if strings.HasPrefix(ipStr, "192.") || strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "172.") {
				return ipStr
			}
		}
	}

	return "localhost"
}
