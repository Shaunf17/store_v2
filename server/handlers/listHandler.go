package handlers

import (
	"encoding/json"
	"net/http"
	"store/auth"
	"store/store"
	"store/utils/logger"

	"github.com/gorilla/mux"
)

type ListItem struct {
	Key   string    `json:"key"`
	Owner auth.User `json:"owner"`
}

func ListHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]

		logger.InfoLogger.Println("You made it to list")
		w.Write([]byte("Welcome to list"))

		if key != "" {
			logger.InfoLogger.Println("Key is", key)
		}
	}
}

func ListGetAll() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var list = []ListItem{}
		s := store.Connect()
		items, err := s.GetAll()
		if err != nil {
			// error
		}

		for _, item := range *items {
			list = append(list, ListItem{
				Key:   item.Key,
				Owner: *item.Owner,
			})
		}

		jsonData, err := json.Marshal(list)
		if err != nil {
			// json error
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func ListGetDetails() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]

		_ = key
	}
}
