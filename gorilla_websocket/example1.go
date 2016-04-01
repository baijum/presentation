package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// START UPGRADER OMIT
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// END UPGRADER OMIT

// END READER OMIT
// START WRITER OMIT
func writer(ws *websocket.Conn) {
	defer ws.Close()

	for {
		ws.WriteMessage(websocket.TextMessage, []byte("hello"))
		time.Sleep(2 * time.Second)
	}
}

// END WRITER OMIT
// START SERVE OMIT
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	go writer(ws)
}

func main() {
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8081", nil)
}

// END SERVE OMIT
