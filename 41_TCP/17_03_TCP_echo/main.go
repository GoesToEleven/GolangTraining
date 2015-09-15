package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("listen messed up", err.Error())
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("conn messed up", err.Error())
			panic(err)
		}

		for {
			var bs = make([]byte, 1024)
			n, err := conn.Read(bs)
			if err != nil {
				break
			}

			m, err := conn.Write(bs)
			if err != nil {
				break
			}
			io.WriteString(conn, fmt.Sprintln("local addr (laddr): ", conn.LocalAddr()))
			io.WriteString(conn, fmt.Sprintln("Remote addr (laddr): ", conn.RemoteAddr()))
			io.WriteString(conn, fmt.Sprintln("Bytes read: ", n))
			io.WriteString(conn, fmt.Sprintln("Bytes written: ", m))

		}
		conn.Close()
	}
}
