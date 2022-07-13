package handlers

import (
	"net/http"
	"store/utils/logger"
)

func ShutdownHandler(shutdown chan<- bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.InfoLogger.Println("Sending shutdown request")
		w.Write([]byte("Shutting down server"))
		shutdown <- true
	}
}
