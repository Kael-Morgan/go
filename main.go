package main

import (
	"context"
	"fmt"
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

	err := http.ListenAndServe(":443", router)
	fmt.Println(err)
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws/{name}", addCORS(websocket_server.ClientHandler))

	mux.HandleFunc("POST /api/v1/cart/{name}", addCORS(handlers.HandleUpdateCartItem))

	mux.HandleFunc("GET /api/v1/cart/{name}", addCORS(handlers.HandleGetCartItems))

	mux.HandleFunc("DELETE /api/v1/cart/{name}", addCORS(handlers.HandleDeleteCartItem))

}

func addCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w.Header())

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Header,Authorization,Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")

		fmt.Println(w.Header(), r)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}
