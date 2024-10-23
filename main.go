package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const PORT string = "localhost:6969";

func main(){
    l, err := net.Listen("tcp4", PORT);
    if err!=nil{
        log.Fatal("LISTEN: ", err);
    }
    defer l.Close();
    for {
        c, err := l.Accept();
        if err!=nil{
            log.Fatal("cannot accept connection on port: ", PORT);
        }
        go handleConnection(c);
    }
}

func handleConnection(c net.Conn){
    fmt.Printf("Serving on %s\n", c.RemoteAddr().String());
    packet := make([]byte, 4096);
    tmp := make([]byte, 4096);
    defer c.Close();
    for {
        _, err := c.Read(tmp);
        if err!=nil{
            if err!=io.EOF{
                log.Fatal("ERROR: ", err)
            }
            println("END OF FILE");
            break
        }
        packet = append(packet, tmp...);
    }
    num, _ := c.Write(packet);
    log.Printf("wrote back %d bytes, the payload is %s\n", num, string(packet));
}
