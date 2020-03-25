package common

import (
	"net"
	"strings"
)

func Telnet(addr, cmd string) (string, error) {
	var answ string
	// connect to unix socket
	c, err := net.Dial("unix", addr)
	if err != nil {
		return answ, err
	}

	// send command to unix socket
	_, err = c.Write([]byte(cmd + "\n"))
	if err != nil {
		return answ, err
	}

	// get answer from unix socket
	b := make([]byte, 512)
	n, err := c.Read(b)
	if err != nil {
		return answ, err
	}

	str := strings.TrimRight(string(b[:n]), " \r\n")

	return str, nil
}
