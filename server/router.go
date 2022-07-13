package server

import (
	"net/http"
	"store/server/handlers"

	"github.com/gorilla/mux"
)

type Handlers map[string]http.HandlerFunc

func Routes(s *Server) map[string]Handlers {
	routes := map[string]Handlers{
		http.MethodGet: map[string]http.HandlerFunc{
			"/ping":        handlers.PingHandler(),
			"/list":        handlers.ListGetAll(),
			"/list/{key}":  handlers.ListGetDetails(),
			"/store/{key}": handlers.StoreGet(),
			"/shutdown":    handlers.ShutdownHandler(s.Channels.Shutdown),
		},
		http.MethodPut: map[string]http.HandlerFunc{
			"/store/{key}": handlers.StorePut(),
		},
		http.MethodDelete: map[string]http.HandlerFunc{
			"/store/{key}": handlers.StoreDelete(),
		},
	}

	return routes
}

func Router(s *Server) http.Handler {
	router := mux.NewRouter()

	for method, handlers := range Routes(s) {
		for p, h := range handlers {
			router.HandleFunc(p, h).Methods(method)
		}
	}

	return router
}
