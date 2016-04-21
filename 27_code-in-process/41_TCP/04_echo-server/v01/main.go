package main

import "net"

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

		for {
			bs := make([]byte, 1024)
			n, err := conn.Read(bs)
			if err != nil {
				break
			}
			_, err = conn.Write(bs[:n])
			if err != nil {
				break
			}
		}

		conn.Close()
	}
}
