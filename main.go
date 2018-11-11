package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const (
	errUnknown   = "Unknown error"
	connTimeout  = "closed"
	connFiltered = "filtered"
	resolveError = "Resolve error"
	lookupError  = "Lookup error"
	connSuccess  = "open"
)

func tcpConnect(address string, port int, timeout time.Duration) (string, error) {
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

func main() {
	addr := "google.com"
	port := 443
	status, err := tcpConnect(addr, port, time.Duration(3)*time.Second)
	format := "%s:%d - %s\n"
	if status == errUnknown {
		fmt.Printf(format, addr, port, err)
	} else {
		fmt.Printf(format, addr, port, status)
	}
}
