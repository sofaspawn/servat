package main

import (
	"log"
	"net"
)

const message = "imp message to be sent"

func main() {
	servAddr := "localhost:6969"
	addr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalln(err)
	}
	reply := make([]byte, 4096)
	_, err = conn.Read(reply)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("> %s\n", string(reply))
}
