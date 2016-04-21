package main

import (
	"bufio"
	"io"
	"net"
	"strings"
)

type databaseRequest struct {
	requestType, key, value string
	resultChannel           chan<- string
}

func handleDatabase(requestChannel <-chan databaseRequest) {
	var db = map[string]string{}
	for request := range requestChannel {
		switch request.requestType {
		case "GET":
			request.resultChannel <- db[request.key]
		case "SET":
			db[request.key] = request.value
			request.resultChannel <- "Set " + request.value + " to key " + request.key + " successful"
		case "DEL":
			delete(db, request.key)
			request.resultChannel <- "Delete of key " + request.key + " successful"
		default:
			request.resultChannel <- "Unknown command: " + request.requestType
		}
		close(request.resultChannel)
	}
}

func handleConn(conn net.Conn, requestChannel chan<- databaseRequest) {
	defer conn.Close()

	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 0 {
			continue
		} else if len(line) < 4 {
			io.WriteString(conn, "Unknown command: "+line+"\n")
			continue
		}
		requestValues := line[4:]
		payload := strings.SplitN(requestValues, " ", 2)
		resultChannel := make(chan string)
		request := databaseRequest{
			requestType:   line[:3],
			key:           payload[0],
			resultChannel: resultChannel,
		}

		if len(payload) > 1 {
			request.value = payload[1]
		}

		requestChannel <- request
		data := <-resultChannel
		io.WriteString(conn, data+"\n")
	}
}

func main() {
	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer server.Close()

	requestChannel := make(chan databaseRequest)
	go handleDatabase(requestChannel)

	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn, requestChannel)
	}
}
