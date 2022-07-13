package server

import (
	"context"
	"net/http"
	"store/configs"
	"store/utils/logger"
)

type Channels struct {
	Shutdown chan bool
	Done     chan bool
}

type Server struct {
	HTTPServer *http.Server
	Router     http.Handler
	Channels   Channels
}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//s.Router.Route(resp, req)
}

func Start(cfg configs.ServerConfigs) error {
	logger.InfoLogger.Println("Starting server...")

	server := &Server{
		HTTPServer: &http.Server{
			Addr: cfg.Port,
		},
		Channels: Channels{
			Shutdown: make(chan bool),
			Done:     make(chan bool),
		},
	}
	server.Router = Router(server)
	server.HTTPServer.Handler = server.Router

	go func() {
		server.HTTPServer.ListenAndServe()
	}()

	logger.InfoLogger.Println("Server started on port", cfg.Port)

	<-server.Channels.Shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger.InfoLogger.Println("Shutting down server...")
	server.HTTPServer.Shutdown(ctx)

	return nil
}
