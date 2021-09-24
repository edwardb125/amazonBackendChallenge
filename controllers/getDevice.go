package controllers

import (
	"amazonBackendChallenge/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetDevice(w http.ResponseWriter, r *http.Request) {
	// connect to dynamoDB in this lines
	db, err := ConnectDynamoDB()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		final, _ := json.Marshal(Error{
			Message: "server error",
		})
		_, _ = w.Write(final)
		return
	}

	vars := mux.Vars(r)
	service := &service.GetCore{
		Db: db,
	}
	item, err := service.GetDevice(vars["id"])

	if err != nil {
		if err.Error() == "server error" {
			w.WriteHeader(500)
			final, _ := json.Marshal(Error{
				Message: "server error",
			})
			_, _ = w.Write(final)
		} else {
			w.WriteHeader(404)
			final, _ := json.Marshal(Error{
				Message: "device not found",
			})
			_, _ = w.Write(final)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	result, _ := json.Marshal(item)
	_, _ = w.Write(result)
}