package main

import (
	"context"
	handlers "go-beyond/handlers/api"
	websocket_server "go-beyond/server"
	"go-beyond/services"
	"net/http"
)

func main() {
	ctx := context.Background()
	services.InitializeRedisClient(ctx)
	websocket_server.InitializeWebSocketServer(ctx)

	router := http.NewServeMux()
	addRoutes(router)

	http.ListenAndServe(":8080", router)

}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws/{name}", websocket_server.ClientHandler)

	mux.HandleFunc("POST /api/v1/cart/{name}", handlers.HandleUpdateCartItem)

	mux.HandleFunc("GET /api/v1/cart/{name}", handlers.HandleGetCartItems)

	mux.HandleFunc("DELETE /api/v1/cart/{name}", handlers.HandleDeleteCartItem)

}
