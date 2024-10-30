package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

const address string = "localhost:6969"

var (
	chats []string
	mu    sync.Mutex
)

func main() {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("LISTEN: ", err)
	}
	defer l.Close()
	log.Printf("Listening on: %s\n", address)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("cannot accept connection: ", err)
			continue
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Printf("Serving on %s\n", c.RemoteAddr().String())

	for {
		packet := make([]byte, 4096)
		n, err := c.Read(packet)
		if err != nil {
			log.Println("ERROR reading from connection: ", err)
			return
		}
		if n == 0 {
			// Connection closed by the client
			return
		}

		// Process the incoming message
		message := strings.TrimSpace(string(packet[:n]))
		if message != "" {
			mu.Lock()
			chats = append(chats, message)
			mu.Unlock()
		}

		// Send back the list of messages
		mu.Lock()
		response := strings.Join(chats, "\n") + "\n"
		mu.Unlock()
		_, err = c.Write([]byte(response))
		if err != nil {
			log.Println("ERROR writing to connection: ", err)
			return
		}
	}
}

