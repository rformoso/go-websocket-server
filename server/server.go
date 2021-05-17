package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/rformoso/go-websocket-server/consts"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connected main page!")
	fmt.Fprintf(w, consts.Logo)
	fmt.Fprintf(w, "version: 1.0.0")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connected websocket service!")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/ws", wsEndpoint)
}

func Start() {
	fmt.Println(consts.Logo)
	setupRoutes()
	log.Fatal(http.ListenAndServeTLS("0.0.0.0:443", "zebraserver.crt", "zebraserver.key", nil))
	// Oops, this shouldn't be here. But take the hint.
	//log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
