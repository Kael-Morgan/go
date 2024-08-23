package main

import (
	"context"
	"fmt"
	"go-beyond/db"
	handlers "go-beyond/handlers/api"
	websocket_server "go-beyond/server"
	"go-beyond/services"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		os.Exit(1)
	}

	ctx := context.Background()

	initServices(ctx)

	handler := initHandler()

	host := os.Getenv("HOST")
	err = http.ListenAndServe(host, handler)
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

	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           600,
	}).Handler(mux)
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws/{name}", websocket_server.ClientHandler)

	mux.HandleFunc("POST /api/v1/cart/{name}", handlers.HandleUpdateCartItem)

	mux.HandleFunc("GET /api/v1/cart/{name}", handlers.HandleGetCartItems)

	mux.HandleFunc("DELETE /api/v1/cart/{name}", handlers.HandleDeleteCartItem)

}
