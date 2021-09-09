package controllers

import (
	"amazonBackendChallenge/dynamoDB"
	"amazonBackendChallenge/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// connect to dynamoDB
	db, err := GetDynamoDB()
	if err != nil {
		log.Println(err)
		CreateError(w, "server error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	service := service.NewGetService(dynamoDB.NewDeviceDB(db))
	item, err := service.GetDevice(vars["id"])
	if err != nil {
		if err.Error() == "server error" {
			CreateError(w, "internal server error", http.StatusInternalServerError)
		} else {
			CreateError(w, "device not found", http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(item)
	_, _ = w.Write(result)
}