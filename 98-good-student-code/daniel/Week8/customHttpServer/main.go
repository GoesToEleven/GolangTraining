package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	// var url string
	headers := map[string]string{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			fields := strings.Fields(line)
			method := fields[0]
			if method != "GET" {
				break
			}
			// url = fields[1]
		} else {
			if line == "" {
				break
			}
			data := strings.SplitN(line, ": ", 2)
			headers[data[0]] = data[1]
		}
	}
	body := `<!DOCTYPE html>
<html>
<head></head>
<body>
  <form method="POST">
    <input type="text" name="KEY">
    <input type="submit">
  </form>
</body>
</html>`

	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)+2)
	io.WriteString(conn, "\r\n")
	fmt.Fprintf(conn, "%s\r\n", body)
}

func main() {
	lnr, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer lnr.Close()

	for {
		conn, err := lnr.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}
