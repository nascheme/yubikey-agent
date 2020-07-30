// +build windows

package main

import (
	"gopkg.in/natefinch/npipe.v2"
	"net"
)

func CreateAuthSock(socketPath string) (net.Listener, error) {
	return npipe.Listen(socketPath)
}
