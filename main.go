package main

import (
	"context"
	"go-beyond/server"
	"go-beyond/services"
	"net/http"
)

func main() {
	ctx := context.Background()
	services.InitializeRedisClient(ctx)
	server.InitializeServer(ctx)

	s := server.GetServer()

	http.ListenAndServe(":8080", &s.Mux)

}
