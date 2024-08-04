package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Server struct {
	clients map[*websocket.Conn]*Client
	mutex   sync.Mutex
	Mux     http.ServeMux
}

type Client struct {
	cartname string
}

var server *Server

func InitializeServer(ctx context.Context) {
	server = &Server{
		clients: make(map[*websocket.Conn]*Client),
		Mux:     *http.NewServeMux(),
	}
	server.addRoutes()
}

func (s *Server) ClientHandler(w http.ResponseWriter, r *http.Request) {
	err := s.handle(w, r)
	if err != nil {
		fmt.Println("Handle error: ", err)
	}
}

func (s *Server) addClient(connection *websocket.Conn, client *Client) {
	s.mutex.Lock()
	s.clients[connection] = client
	s.mutex.Unlock()
	fmt.Println("Added Client ", connection)
}

func (s *Server) handle(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	cartName := r.PathValue("name")
	if cartName == "" {
		return nil
	}
	client := &Client{
		cartname: r.PathValue("name"),
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

	for {
		var message string
		err = wsjson.Read(ctx, conn, &message)
		if err != nil {
			return err
		}

		fmt.Println("Received message: ", message)
	}

}
func GetServer() *Server {
	return server
}
