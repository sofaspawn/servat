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
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        log.Fatalln("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()
    log.Printf("Listening on %s:%s\n", CONN_HOST, CONN_PORT)
    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatalln("Error accepting: ", err.Error())
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    log.Printf("handling connection: %s\n", conn.RemoteAddr().String())
    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        log.Fatalln("Error reading:", err.Error())
    }
    buf = append(buf, byte('\n'))
    //conn.Write([]byte("Message received."))
    conn.Write(buf)
    conn.Close()
}
