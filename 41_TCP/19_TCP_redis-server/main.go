package main

import (
	"log"
	"net"
	"fmt"
)

var db = make(map[string]string)

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

		go func() {
			for {
				var bs = make([]byte, 1024)
				n, err := conn.Read(bs)
				if err != nil {
					break
				}

				// check bs to see if first word is GET, SET, DEL
				// respond appropriately
				fmt.Println(bs[:n])
				fmt.Println(string(bs[:n]))

				_, err = conn.Write(bs[:n])
				if err != nil {
					break
				}

			}
			conn.Close()
		}()

	}

}
