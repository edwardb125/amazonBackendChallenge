package controllers

import (
	//"./createDynamoDB"
	"amazonBackendChallenge/models"
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetDeviceController(t *testing.T) {
	_ = os.Unsetenv("AWS_REGION")
	_ = os.Unsetenv("TABLE_NAME")
	input := models.Device{
		Id:          "333",
		DeviceModel: "a",
		Name:        "a",
		Note:        "a",
		Serial:      "a",
	}
	tests := []struct {
		name   string
		input  models.Device
		status int
		output interface{}
	}{
		{name: "invalid input", input: models.Device{
			Id: "1",
		}, status: 400, output: Error{
			Message: "invalid device info",
		}},
		{name: "server error", input: input, status: 500, output: Error{
			Message: "server error",
		}},
		{name: "ok", input: input, status: 201, output: input},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "ok" {
				_ = os.Setenv("AWS_REGION", "us-west-2")
				_ = os.Setenv("TABLE_NAME", "Devices")
			}
			router := mux.NewRouter()
			router.HandleFunc("/devices", SetDevice).Methods("POST")

			marshal, _ := json.Marshal(test.input)
			req, _ := http.NewRequest(http.MethodPost, "/devices", bytes.NewBuffer(marshal))

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			assert.Equal(t, test.status, rr.Code)
			if rr.Code == 201 {
				var device models.Device
				_ = json.Unmarshal(rr.Body.Bytes(), &device)
				assert.Equal(t, test.output.(models.Device), device)
			} else {
				var err Error
				_ = json.Unmarshal(rr.Body.Bytes(), &err)
				assert.Equal(t, test.output.(Error), err)
			}

		})
	}
	//clear data in dynamoDB
	DeleteItem(t, input.Id)

}