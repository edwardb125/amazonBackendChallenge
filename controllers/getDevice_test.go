package controllers

import (
	"amazonBackendChallenge/models"
	"amazonBackendChallenge/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	input := models.Device{
		Id:          "ididid",
		DeviceModel: "test",
		Name:        "test",
		Note:        "test",
		Serial:      "test",
	}

	_ = os.Setenv("AWS_REGION", "us-west-1")
	_ = os.Setenv("TABLE_NAME", "Devices")
	db, err := ConnectDynamoDB()
	if err != nil{
		t.Fatal("error occurred while connecting to dynamodb")
	}
	temp := &service.CreateCore{
		Db: db,
	}
	err = temp.CreateDevice(input)
	if err != nil {
		t.Fatal("error occurred while device creating", err)
	}
	_ = os.Unsetenv("AWS_REGION")
	_ = os.Unsetenv("TABLE_NAME")

	tests := []struct {
		name   string
		id     string
		status int
		output interface{}
	}{
		{name: "server error", status: 500, output: Error{
			Message: "server error",
		}, id: "ididid"},
		{name: "not found error", status: 404, output: Error{
			Message: "device not found",
		}, id: "wrongId"},
		{name: "well done", status: 200, output: input, id: input.Id},

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name != "server error" {
				_ = os.Setenv("AWS_REGION", "us-west-1")
				_ = os.Setenv("TABLE_NAME", "Devices")
			} else{
				_ = os.Unsetenv("AWS_REGION")
				_ = os.Setenv("TABLE_NAME", "Devices")
				log.Println(os.Getenv("AWS_REGION"))
			}

			router := mux.NewRouter()
			router.HandleFunc("/devices/{id}", GetDevice).Methods("GET")
			req, _ := http.NewRequest(http.MethodGet, "/devices/"+test.id, nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, test.status, res.Code)

			if test.status == 200 {
				var device models.Device
				_ = json.Unmarshal(res.Body.Bytes(), &device)
				assert.Equal(t, test.output.(models.Device), device)
			} else {
				var message Error
				_ = json.Unmarshal(res.Body.Bytes(), &message)
				assert.Equal(t, test.output.(Error), message)
			}
		})
	}
	DeleteDeviceId(t, input.Id)
}