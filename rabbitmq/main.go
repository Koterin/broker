package main

import (
    "log"
    "net/http"

    "github.com/koterin/broker/rabbitmq/server"
)

func main() {
    log.SetPrefix("[LOG] ")
    log.SetFlags(3)

    log.Printf("Server started successfully")

    http.HandleFunc("/sendNote", server.SendNote)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
