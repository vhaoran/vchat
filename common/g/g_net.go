package g

import (
	"net"
)

func IpAddrOfLan() []string {
	l := make([]string, 0)

	a, _ := net.InterfaceAddrs()
	for _, address := range a {
		// 检查ip地址判断是否回环地址
		if ip, ok := address.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				l = append(l, ip.IP.String())
			}
		}
	}
	return l
}
