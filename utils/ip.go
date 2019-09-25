package utils

import "net"

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

			if ipnet.IP.To4() != nil {
				//log.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
