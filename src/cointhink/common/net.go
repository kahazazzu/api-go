package common

import (
	"cointhink/httpclients"
	"cointhink/q"

	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Httploop(listen_address string) {
	http.HandleFunc("/", Upgrade)
	httpclients.Clients = map[*websocket.Conn]httpclients.Httpclient{}
	http.ListenAndServe(listen_address, nil)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s webstocket started", r.Header.Get("Origin"))
	wsocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("websocket upgrade fail:", err)
		return
	}
	_client := httpclients.Httpclient{Socket: wsocket, Out: []*q.RpcOut{}}
	log.Printf("wsocket open %p", wsocket)
	httpclients.Clients[wsocket] = _client
	for {
		_, payload, err := wsocket.ReadMessage()
		if err != nil {
			log.Println("ws_recv err:", err)
			break
		}
		RPCq <- q.RpcMsg{Socket: _client.Socket, Payload: payload}
	}
	log.Printf("wsocket closing %p", wsocket)
	wsocket.Close()
	delete(httpclients.Clients, wsocket)
}
