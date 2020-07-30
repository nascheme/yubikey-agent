// +build !windows

package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

func CreateAuthSock(socketPath string) (net.Listener, error) {
	os.Remove(socketPath)
	if err := os.MkdirAll(filepath.Dir(socketPath), 0777); err != nil {
		log.Fatalln("Failed to create UNIX socket folder:", err)
	}
	return net.Listen("unix", socketPath)
}
