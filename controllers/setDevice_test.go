package controllers

import (
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
		Id:          "ididid",
		DeviceModel: "test",
		Name:        "test",
		Note:        "test",
		Serial:      "test",
	}
	tests := []struct {
		name   string
		input  models.Device
		status int
		output interface{}
	}{
		{name: "server error", input: input, status: 500, output: Error{
			Message: "server error",
		}},
		{name: "input is invalid", input: models.Device{
			Id: "1",
		}, status: 400, output: Error{
			Message: "invalid device info",
		}},
		{name: "well done", input: input, status: 201, output: input},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name != "server error" {
				_ = os.Setenv("AWS_REGION", "us-west-1")
				_ = os.Setenv("TABLE_NAME", "Devices")
			}else{
				_ = os.Unsetenv("AWS_REGION")
				_ = os.Setenv("TABLE_NAME", "Devices")
			}
			router := mux.NewRouter()
			router.HandleFunc("/devices", SetDevice).Methods("POST")

			marshal, _ := json.Marshal(test.input)
			req, _ := http.NewRequest(http.MethodPost, "/devices", bytes.NewBuffer(marshal))

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, test.status, res.Code)

			if res.Code == 201 {
				var device models.Device
				_ = json.Unmarshal(res.Body.Bytes(), &device)
				assert.Equal(t, test.output.(models.Device), device)
			} else {
				var err Error
				_ = json.Unmarshal(res.Body.Bytes(), &err)
				assert.Equal(t, test.output.(Error), err)
			}
		})
	}
	DeleteDeviceId(t, input.Id)

}