package main

import (
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
