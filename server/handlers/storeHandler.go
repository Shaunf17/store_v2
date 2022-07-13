package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"store/auth"
	"store/store"
	"store/utils/logger"

	"github.com/gorilla/mux"
)

var s = store.Connect()

func StoreGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]

		entity, err := s.Find(key)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(entity)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(jsonData)
	}
}

func StorePut() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		msg, err := s.Add(key, string(body), auth.BasicAuth(r))
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(msg))
	}
}

func StoreDelete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]

		_, err := s.Find(key)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		msg, err := s.Delete(key)
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(msg))
	}
}
