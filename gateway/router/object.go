package router

import (
	"shaos/gateway/object"

	"github.com/gorilla/mux"
)

func ObjectRoutes(s *mux.Router) {
	s.HandleFunc("/{name}", object.GetHandler).Methods("Get")
	s.HandleFunc("/{name}", object.PutHandler).Methods("PUT")
}
