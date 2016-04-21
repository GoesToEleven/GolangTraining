package main

import (
	"io"
	"log"
	"net"
)

type Rot13Reader struct {
	io.Reader
}

func (rot13 *Rot13Reader) Read(p []byte) (int, error) {
	n, err := rot13.Reader.Read(p)
	for i, v := range p[:n] {
		if v <= 'z' && v >= 'a' {
			p[i] = v + 13
			if p[i] > 'z' {
				p[i] -= 26
			}
		} else if v <= 'Z' && v >= 'A' {
			p[i] = v + 13
			if p[i] > 'Z' {
				p[i] -= 26
			}
		} else {
			p[i] = v
		}
	}
	return n, err
}

func rot13(rdr io.Reader) io.Reader {
	return &Rot13Reader{rdr}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, rot13(conn))
}

func main() {
	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleConn(conn)
	}
}
