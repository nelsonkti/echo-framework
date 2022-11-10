package helper

import (
	"fmt"
	"github.com/nelsonkti/echo-framework/lib/logger"
	"net"
	"runtime"
)

//获取本机ip
var ip string

func GetLocalIP() string {
	if ip != "" {
		return ip
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return ip
			}
		}
	}
	return ""
}

// RecoverPanic 恢复panic
func RecoverPanic() {
	err := recover()
	if err != nil {
		logger.Sugar.Error(err)

		buf := make([]byte, 2048)
		n := runtime.Stack(buf, false)
		logger.Sugar.Error(fmt.Sprintf("%s", buf[:n]))
	}
}
