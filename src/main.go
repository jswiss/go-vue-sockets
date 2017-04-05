package main

import(
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)  // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{}

// Define our message object
type Message struct {
        Email    string `json:"email"`
        Username string `json:"username"`
        Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages concurrently
	go handleMessages()

	// Start webserver
	log.Println("http server started on localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
