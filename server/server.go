package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    //small change
)

const port = ":8000"

func main() {

    //chats := make([]string, 1024)

    router := mux.NewRouter()

    fmt.Printf("serving on http://localhost%s\n", port)

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprint(w, "Hello World")
    })

    router.HandleFunc("/chats", handleChat).Methods(http.MethodGet)

    http.ListenAndServe(port, router)
}

func getChats() []string {
    chats := []string{"hello sofa!\n", "yes?\n", "how r u?\n"}
    return chats
}

func handleChat(w http.ResponseWriter, r *http.Request){
    chats := getChats()
    for _, chat := range chats{
        fmt.Fprint(w, chat)
    }
}
