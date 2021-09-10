package controllers

import (
	"amazonBackendChallenge/dynamoDB"
	"amazonBackendChallenge/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func SetDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// get data from request body and change it to object with device attribute
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)

	//validate data with device pattern
	validate := validator.New()
	err := validate.Struct(device)
	if err != nil {
		log.Println(err)
		dynamoDB.CreateError(w, "invalid device attribute", http.StatusBadRequest)
		return
	}

	dynamoDB.DoWithDynamoDB(w,device)
}
