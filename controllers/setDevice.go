package controllers

import (
	"amazonBackendChallenge/models"
	"amazonBackendChallenge/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func SetDevice(w http.ResponseWriter, r *http.Request) {
	// get data from request body and change it to object with device attribute
	var device models.Device
	_ = json.NewDecoder(r.Body).Decode(&device)

	//validate data with device pattern
	validate := validator.New()
	err := validate.Struct(device)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		final, _ := json.Marshal(Error{
			Message: "invalid device attribute",
		})
		_, _ = w.Write(final)
		return
	}

	//connect to the dynamoDB
	db, err := ConnectDynamoDB() // this function should be created
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		final, _ := json.Marshal(Error{
			Message: "server error",
		})
		_, _ = w.Write(final)
		return
	}

	//send data to dynamoDB
	service := &service.CreateCore{
		Db: db,
	}
	err = service.CreateDevice(device)

	if err != nil{
		log.Println(err)
		w.WriteHeader(500)
		final, _ := json.Marshal(Error{
			Message: "server error",
		})
		_, _ = w.Write(final)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	final, _ := json.Marshal(device)
	_, _ = w.Write(final)
}

