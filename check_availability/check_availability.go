package check_availability

import (
	"net"
	"strconv"
	"time"
)

func Check_Service_Availability(target string, port int) bool {
	//fmt.Print("Check if Target Service is available: ")
	checkTimeout := 5
	serverAddress := net.JoinHostPort(target, strconv.Itoa(port))
	timeout := time.Second * time.Duration(checkTimeout)
	tcpConn, tcpErr := net.DialTimeout("tcp", serverAddress, timeout)
	tcpResult := "FAIL"
	if tcpErr == nil {
		tcpResult = "OK"
	}
	if tcpResult == "OK" {
		tcpConn.Close()
		//fmt.Println("OK")
		//fmt.Println("")
		return true
	} else {
		//fmt.Println("FAIL")
		return false
	}
}
