package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const (
	errUnknown   = -5
	connTimeout  = -4
	connFiltered = -3
	resolveError = -2
	lookupError  = -1
	connSuccess  = 1
)

func tcpConnect(address string, port int, timeout time.Duration) (int, error) {
	ipList, err := net.LookupHost(address)
	if err != nil {
		return lookupError, err
	}
	if len(ipList) == 0 {
		return resolveError, fmt.Errorf("Unable to resolve %s", address)
	}
	connAddr := fmt.Sprintf("%s:%d", ipList[0], port)
	conn, err := net.DialTimeout("tcp", connAddr, timeout)
	if err != nil {
		sErr := err.Error()
		if strings.HasSuffix(sErr, "connection refused") {
			return connFiltered, err
		}
		if strings.HasSuffix(sErr, "i/o timeout") {
			return connTimeout, err
		}
		return errUnknown, err
	}
	defer conn.Close()

	return connSuccess, nil
}

func statusToStr(status int) string {
	strStatus := map[int]string{
		errUnknown:   "Unknown error",
		connTimeout:  "closed",
		connFiltered: "filtered",
		resolveError: "Resolve error",
		lookupError:  "Lookup error",
		connSuccess:  "open",
	}
	return strStatus[status]
}

func main() {
	addr := "google.com"
	port := 443
	statusInt, err := tcpConnect(addr, port, time.Duration(3)*time.Second)
	format := "%s:%d - %s\n"
	if statusInt == errUnknown {
		fmt.Printf(format, addr, port, err)
	} else {
		fmt.Printf(format, addr, port, statusToStr(statusInt))
	}
}
