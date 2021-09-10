package dynamoDB

import (
	"amazonBackendChallenge/models"
	service2 "amazonBackendChallenge/service"
	"encoding/json"
	"log"
	"net/http"
)


func DoWithDynamoDB(w http.ResponseWriter, device models.Device){

	//connect to the dynamoDB
	db, err := GetDynamoDB() // this function should be created
	if err != nil {
		log.Println(err)
		CreateError(w, "server error", http.StatusInternalServerError)
		return
	}

	//send data to dynamoDB
	service := service2.NewCreateService(NewDeviceDB(db))
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
