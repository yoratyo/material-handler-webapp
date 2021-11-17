package shared

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"time"
)

func GetCurrentTime() (string, string) {
	now := time.Now()
	y, m, d := now.Date()
	currDate := fmt.Sprintf("%d-%d-%d", y, int(m), d)
	currTime := fmt.Sprintf("%d:%d:%d", now.Hour(), now.Minute(), now.Second())

	return currDate, currTime
}

func GetCurrentInMilis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func SetErrorCookie(c *gin.Context, msg string) {
	c.SetCookie("error", msg, 10, "/", c.Request.URL.Hostname(), false, true)
	c.SetCookie("errorTime", strconv.FormatInt(GetCurrentInMilis(), 10), 10, "/", c.Request.URL.Hostname(), false, true)
}

func SetSuccessCookie(c *gin.Context, msg string) {
	c.SetCookie("success", msg, 10, "/", c.Request.URL.Hostname(), false, true)
	c.SetCookie("successTime", strconv.FormatInt(GetCurrentInMilis(), 10), 10, "/", c.Request.URL.Hostname(), false, true)
}

func RemoveSuccessCookie(c *gin.Context) {
	c.SetCookie("success", "", 0, "/", c.Request.URL.Hostname(), false, true)
	c.SetCookie("successTime", "", 0, "/", c.Request.URL.Hostname(), false, true)
}

func RemoveErrorCookie(c *gin.Context) {
	c.SetCookie("error", "", 0, "/", c.Request.URL.Hostname(), false, true)
	c.SetCookie("errorTime", "", 0, "/", c.Request.URL.Hostname(), false, true)
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
				ip = ipnet.IP.String()
			}
		}
	}

	return ip, nil
}
