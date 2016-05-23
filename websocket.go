package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/websocket"
)

func echo(ws *websocket.Conn) {
	/*msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:n])

	m, err := ws.Write(msg[:n])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", msg[:m])*/
	fmt.Printf("Got a connection")
	io.Copy(ws, ws)
}
func websocketHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	websocket.Handler(echo).ServeHTTP(w, r)
}
