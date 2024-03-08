package livehost

import (
	"fmt"
	"net"
)

func SideroLiveHost(i int) (net.IP, error) {
	address := fmt.Sprintf("192.168.1.%d:50000", i)
	ip := fmt.Sprintf("192.168.1.%d", i)
	var LiveHost net.IP
	conn, err := net.Dial("tcp", address)
	if err == nil {
		conn.Close()
		LiveHost = net.ParseIP(ip)
		return LiveHost, nil
	} else {
		return nil, nil
	}
}
