package main

import (
<<<<<<< HEAD
	"log"
	"net"
	"os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "6969"
    CONN_TYPE = "tcp"
)

func main() {
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT);
    if err!=nil{
        log.Fatalln(err);
        os.Exit(1);
    }
    defer l.Close();
    for{
        conn, err := l.Accept();
        if err!=nil{
            log.Fatalln(err);
        }
        buf := make([]byte, 1024);
        conn.Read(buf);
    }
}

func handleRequest
=======
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

>>>>>>> 10fe41b2474e37fc902a6ee356a6db5704bab365
