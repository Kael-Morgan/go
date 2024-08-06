package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
)

type Server struct {
	clients map[*websocket.Conn]*Client
	Mutex   sync.Mutex
}

type Client struct {
	CartName string
}

var server *Server

func InitializeWebSocketServer(ctx context.Context) {
	server = &Server{
		clients: make(map[*websocket.Conn]*Client),
	}
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	err := server.handle(w, r)
	if err != nil {
		fmt.Println("Handle error: ", err)
	}
}

func (s *Server) addClient(connection *websocket.Conn, client *Client) {
	s.Mutex.Lock()
	s.clients[connection] = client
	s.Mutex.Unlock()

	fmt.Println(len(s.clients))
}

func (s *Server) handle(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	cartName := r.PathValue("name")
	if cartName == "" {
		return nil
	}
	client := &Client{
		CartName: r.PathValue("name"),
	}

	opts := &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	}

	conn, err := websocket.Accept(w, r, opts)

	s.addClient(conn, client)
	if err != nil {
		return err
	}

	defer conn.Close(websocket.StatusInternalError, "Internal error")
	defer func() { delete(s.clients, conn) }()

	for {
		_, message, err := conn.Read(ctx)
		fmt.Println(string(message))
		if err != nil {
			return err
		}

		fmt.Println("Received message: ", message)
	}

}
func GetServer() *Server {
	return server
}

func GetClients() map[*websocket.Conn]*Client {
	return server.clients
}
