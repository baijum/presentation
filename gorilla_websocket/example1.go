package main

// START IMPORT OMIT
import (
	"net/http"
	"time"

	"github.com/gorilla/websocket" // HL
)

// END IMPORT OMIT

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
		ws.WriteMessage(websocket.TextMessage, []byte("hello")) // HL
		time.Sleep(2 * time.Second)
	}
}

// END WRITER OMIT

// START SERVE OMIT
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil) // HL
	go writer(ws)
}

// END SERVE OMIT

// START MAIN OMIT
func main() {
	http.HandleFunc("/ws", serveWs) // HL
	http.ListenAndServe(":8081", nil)
}

// END MAIN OMIT
