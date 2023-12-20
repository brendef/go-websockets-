package main

import (
	"net/http"
	"websockets/websockets"

	websocketPkg "golang.org/x/net/websocket"
)

func main() {

	ws := websockets.NewWebSocket()
	http.Handle("/ws", websocketPkg.Handler(ws.HandleWebSocket))

	http.ListenAndServe(":8080", nil)

}
