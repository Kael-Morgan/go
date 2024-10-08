package main

import (
	"context"
	"fmt"
	"go-beyond/db"
	handlers "go-beyond/handlers/api"
	websocket_server "go-beyond/server"
	"go-beyond/services"
	"net/http"
)

func main() {
	ctx := context.Background()

	initServices(ctx)

	handler := initHandler()

	err := http.ListenAndServe(":8080", handler)
	fmt.Println(err)
}

func initServices(ctx context.Context) {
	services.InitializeRedisClient(ctx)
	db.InitializeDB(ctx)
	websocket_server.InitializeWebSocketServer(ctx)
}

func initHandler() http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)

	// return cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{
	// 		http.MethodHead,
	// 		http.MethodGet,
	// 		http.MethodPost,
	// 		http.MethodPut,
	// 		http.MethodPatch,
	// 		http.MethodDelete,
	// 	},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: false,
	// 	MaxAge:           600,
	// }).Handler(mux)
	return mux
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws/{name}", websocket_server.ClientHandler)

	mux.HandleFunc("POST /api/v1/cart/{name}", handlers.HandleUpdateCartItem)

	mux.HandleFunc("GET /api/v1/cart/{name}", handlers.HandleGetCartItems)

	mux.HandleFunc("DELETE /api/v1/cart/{name}", handlers.HandleDeleteCartItem)

}
