package main

import (
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// handles only one connection
		io.Copy(conn, conn)
		conn.Close()
	}
}

/*
start main.go (go run main.go) then ...
telnet localhost 9000
*/
