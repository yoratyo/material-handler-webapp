package shared

import (
	"fmt"
	"net"
	"time"
)

func GetCurrentTime() (string, string) {
	now := time.Now()
	y, m, d := now.Date()
	currDate := fmt.Sprintf("%d-%d-%d", y, int(m), d)
	currTime := fmt.Sprintf("%d:%d:%d", now.Hour(), now.Minute(), now.Second())

	return currDate, currTime
}

func GetServerIPAddress() (string, error) {
	var ip string

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, a := range addrs {

		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(">>>>>>>>>>>> addrs : " + a.String())
				ip = ipnet.IP.String()
			}
		}
	}

	return ip, nil
}
