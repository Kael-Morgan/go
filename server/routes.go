package server

import (
	handlers "go-beyond/handlers/api"
)

func (s *Server) addRoutes() {
	s.Mux.HandleFunc("/ws/{name}", s.ClientHandler)

	s.Mux.HandleFunc("POST /api/v1/cart/{name}", handlers.HandleUpdateCartItem)

	s.Mux.HandleFunc("GET /api/v1/cart/{name}", handlers.HandleGetCartItems)

	s.Mux.HandleFunc("DELETE /api/v1/cart/{name}", handlers.HandleDeleteCartItem)

}
