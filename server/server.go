package main;

import (
	"fmt"
	//"io"
	"log"
	"net"
)

const address string = "localhost:6969";

func main(){
    chats := make([]byte, 10*1024);

    l, err := net.Listen("tcp4", address);

    if err!=nil{
        log.Fatal("LISTEN: ", err);
    } else {
        log.Printf("Listening on: http://%s\n", address);
    }
    defer l.Close();
    for {
        c, err := l.Accept();
        if err!=nil{
            log.Fatal("cannot accept connection on port: ", address);
        }
        go handleConnection(c, &chats);
    }
}

func handleConnection(c net.Conn, chats *[]byte){
    fmt.Printf("Serving on %s\n", c.RemoteAddr().String());
    packet := make([]byte, 4096);
    //tmp := make([]byte, 4096);
    _, err := c.Read(packet);
    if err!=nil{
        log.Fatal("ERROR: ", err);
    }
    defer c.Close();
    *chats = append(*chats, packet...)
    num, _ := c.Write(*chats);
    log.Printf("wrote back %d bytes\n", num);
}
