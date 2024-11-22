package main

import (
	"log"
	"net"
	"strings"
)

const message = "imp message to be sent"

func handleError(err *error){
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	servAddr := "localhost:6969"
	addr, err := net.ResolveTCPAddr("tcp", servAddr)
	conn, err := net.DialTCP("tcp", nil, addr)

    handleError(&err)

	defer conn.Close()
	_, err = conn.Write([]byte(message))

    handleError(&err)

	reply := make([]byte, 1024*100)
	_, err = conn.Read(reply)

    handleError(&err)

	for _, msg := range strings.Split(string(reply), "\n") {
		log.Printf("> %s", msg)
	}
}
