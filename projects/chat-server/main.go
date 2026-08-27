package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	clients = make(map[*websocket.Conn]bool)
	mu  sync.Mutex
)

func deferCleanUp(conn *websocket.Conn)  {
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	fmt.Printf("Clients connected: %d\n", len(clients))
}

func ws(w http.ResponseWriter, r *http.Request)  {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	defer deferCleanUp(conn)

	fmt.Printf("Clients connected: %d\n", len(clients))

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Received: %s\n", string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", ws)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}