package main

import (
	"context"
	"fmt"
	handlers "go-beyond/handlers/api"
	websocket_server "go-beyond/server"
	"go-beyond/services"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load(".env")

	ctx := context.Background()
	services.InitializeRedisClient(ctx)
	websocket_server.InitializeWebSocketServer(ctx)

	router := http.NewServeMux()

	addRoutes(router)

	corsHandler := cors.AllowAll().Handler(router)

	err = http.ListenAndServe(":443", corsHandler)

	fmt.Println(err)
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws/{name}", websocket_server.ClientHandler)

	mux.HandleFunc("POST /api/v1/cart/{name}", handlers.HandleUpdateCartItem)

	mux.HandleFunc("GET /api/v1/cart/{name}", handlers.HandleGetCartItems)

	mux.HandleFunc("DELETE /api/v1/cart/{name}", handlers.HandleDeleteCartItem)

	// mux.HandleFunc("/ws/{name}", addCORS(websocket_server.ClientHandler))

	// mux.HandleFunc("POST /api/v1/cart/{name}", addCORS(handlers.HandleUpdateCartItem))

	// mux.HandleFunc("GET /api/v1/cart/{name}", addCORS(handlers.HandleGetCartItems))

	// mux.HandleFunc("DELETE /api/v1/cart/{name}", addCORS(handlers.HandleDeleteCartItem))

}

// func addCORS(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Headers", "*")
// 		w.Header().Set("Access-Control-Allow-Credentials", "false")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")

// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	}
// }
