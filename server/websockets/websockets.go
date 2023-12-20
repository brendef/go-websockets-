package websockets

import (
	"encoding/json"
	"io"
	"log"
	"websockets/lib"
	"websockets/models"

	websocketPkg "golang.org/x/net/websocket"
)

type WebSocket struct {
	connections map[*websocketPkg.Conn]bool
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		connections: make(map[*websocketPkg.Conn]bool),
	}
}

func (s *WebSocket) HandleWebSocket(ws *websocketPkg.Conn) {
	log.Println("New connection: ", ws.RemoteAddr())

	s.connections[ws] = true

	buffer := make([]byte, 1024)

	for {
		n, err := ws.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Println(err)
			continue
		}

		msgBytes := buffer[:n]

		var clientReq models.ClientReq
		err = json.Unmarshal(msgBytes, &clientReq)
		if err != nil {
			log.Println("error in unmarshal", err)
		}

		log.Println("subscribing to ", clientReq.Data)

		write := func(ws *websocketPkg.Conn, res map[string]int) {

			// send response to client
			if err := websocketPkg.JSON.Send(ws, res); err != nil {
				log.Println("error in send", err)
			}

			// if _, err := ws.Write([]byte(databaseResponse.ToString())); err != nil {
			// 	log.Println("error in broadcast", err)
			// }
		}

		for ws := range s.connections {
			changesChannel := lib.MonitorFileChanges("database.txt")
			for change := range changesChannel {
				log.Println("File content changed:", change)
				s, err := lib.TextToMap(change)
				if err != nil {
					log.Println("error in text to map", err)
				}

				go write(ws, s)
			}

		}
	}

}
