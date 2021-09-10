package controllers

import (
	"amazonBackendChallenge/dynamoDB"
	"amazonBackendChallenge/models"
	service2 "amazonBackendChallenge/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func SetDevice(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type","application/json")

	// get data from request body and change it to object with device attribute
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)

	//validate data with device pattern
	validate := validator.New()
	err := validate.Struct(device)
	if err != nil {
		log.Println(err)
		CreateError(w,"invalid device attribute", http.StatusBadRequest)
		return
	}

	//connect to the dynamoDB
	db, err := GetDynamoDB() // this function should be created
	if err != nil {
		log.Println(err)
		CreateError(w, "server error", http.StatusInternalServerError)
		return
	}

	//send data to dynamoDB
	service := service2.NewCreateService(dynamoDB.NewDeviceDB(db))
	err = service.CreateDevice(device)
	if err != nil{
		log.Println(err)
		CreateError(w, "server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	result, _ := json.Marshal(device)
	_, _ = w.Write(result)
}