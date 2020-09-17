package test

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client" + reply)

		msg := "Received:" + reply
		fmt.Println("Sending to client:" + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func MyWs() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
