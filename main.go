package main

import (
    "fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request){
    conn, err := upgrader.Upgrade(w, r, nil)
    if err!=nil{
        log.Println(err)
        return
    }
    fmt.Println(conn.ReadMessage())
}

func main(){
}
