package handlers

import (
	"net/http"
	"store/utils/logger"
)

func PingHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.InfoLogger.Printf("Ping called from %v - Replying...\n", r.Host)
		w.Write([]byte("pong"))
	}
}
