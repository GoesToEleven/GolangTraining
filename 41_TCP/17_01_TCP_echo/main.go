package main

import (
	"net"
	"log"
	"io"
	"fmt"
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

		var bs = make([]byte, 1024)
		n, err := conn.Read(bs)
		if err != nil {
			log.Fatalln("read messed up", err.Error())
			panic(err)
		}

		m, err := conn.Write(bs)
		if err != nil {
			log.Fatalln("write messed up", err.Error())
			panic(err)
		}

		io.WriteString(conn, fmt.Sprintln("Bytes read: ", n))
		io.WriteString(conn, fmt.Sprintln("Bytes written: ", m))

		conn.Close()
	}
}
