package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/ttacon/chalk"
)

// User holds a user of the chat room
type User struct {
	Name   string
	Output chan Message
}

// Message holds a message for the chat room
type Message struct {
	Username string
	Text     string
}

// ChatServer holds all the data about the current state of the chat room
type ChatServer struct {
	Users map[string]User
	Join  chan User
	Leave chan User
	Input chan Message
}

// Run runs the server
func (cs *ChatServer) Run() {
	for {
		select {
		case user := <-cs.Join:
			cs.Users[user.Name] = user
			go func() {
				cs.Input <- Message{
					Username: "SYSTEM",
					Text:     fmt.Sprintf("%s joined", user.Name),
				}
			}()
		case user := <-cs.Leave:
			delete(cs.Users, user.Name)
			go func() {
				cs.Input <- Message{
					Username: "SYSTEM",
					Text:     fmt.Sprintf("%s left", user.Name),
				}
			}()
		case msg := <-cs.Input:
			for _, user := range cs.Users {
				select {
				case user.Output <- msg:
				default:
				}
			}
		}
	}
}

func handleConn(chatServer *ChatServer, conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "Enter your username: ")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()

	user := User{
		Name:   scanner.Text(),
		Output: make(chan Message, 10),
	}

	chatServer.Join <- user
	defer func() {
		chatServer.Leave <- user
	}()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			chatServer.Input <- Message{user.Name, ln}
		}
		close(user.Output)
	}()

	for msg := range user.Output {
		_, err := fmt.Fprintf(conn, "%s%s: %s%s%s\n", chalk.Yellow, msg.Username, chalk.White, msg.Text, chalk.ResetColor)
		if err != nil {
			break
		}
	}
}

func main() {
	lnr, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer lnr.Close()

	chatServer := &ChatServer{
		Users: make(map[string]User),
		Join:  make(chan User),
		Leave: make(chan User),
		Input: make(chan Message),
	}

	go chatServer.Run()

	for {
		conn, err := lnr.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(chatServer, conn)
	}
}
