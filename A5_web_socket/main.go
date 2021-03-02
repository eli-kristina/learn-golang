package main

import (
	"net/http"
	"html/template"
	"strings"
	"path"
	"fmt"

	"github.com/gorilla/websocket"
)

type M map[string]interface{}

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

const MESSAGE_NEW = "NEW"
const MESSAGE_CHAT = "CHAT"
const MESSAGE_LEAVE = "LEAVE"

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Web Socket",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)

		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		username := r.URL.Query().Get("username")
		currentConn := WebSocketConnection{Conn: wsConn, Username: username}
		connections = append(connections, &currentConn)

		go handleIO(&currentConn, connections)
	})

	fmt.Println("server started at localhost:9090")
	http.ListenAndServe(":9090", nil)
}

func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, MESSAGE_NEW, "")

	for {
		payload := SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, MESSAGE_LEAVE, "")
				return
			}

			fmt.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, MESSAGE_CHAT, payload.Message)
	}
}

func broadcastMessage(currentConn *WebSocketConnection, kind, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		})
	}
}