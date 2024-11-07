package main;

import (
	"net"
	"log"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:6969");
	if err!=nil{
		log.Fatalln(err);
	}
	defer conn.Close();
	_, err = conn.Write([]byte("hallo"))
	if err!=nil{
		log.Fatalln(err)
	}
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err!=nil{
		log.Fatalln(err);
	}
	if n>0{
		log.Printf("Message: %s\n", string(buffer[:n]))
	} else {
		log.Println("No messages received!")
	}
}
