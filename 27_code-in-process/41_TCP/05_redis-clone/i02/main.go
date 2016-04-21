package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		// skip blank lines
		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
		case "SET":
		case "DEL":
		default:
			io.WriteString(conn, "INVALID COMMAND "+fs[0]+"\n")
		}
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn)
	}
}
