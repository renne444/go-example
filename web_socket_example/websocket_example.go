package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	
	if err != nil {
		w.Write([]byte("fff"))
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		log.Printf("msg: %s", string(msg))
	}
}

func main() {
	http.HandleFunc("/", ws)
	if err := http.ListenAndServe("0.0.0.0:40444", nil); err != nil {
		log.Fatal("listen failed error = " + err.Error())
	}
}
