package main

import (
	"amazonBackendChallenge/controllers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillaMux"
	"github.com/gorilla/mux"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){

	// set router with gorilla mux
	router := mux.NewRouter()
	router.HandleFunc("/devices", controllers.SetDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", controllers.GetDevice).Methods("GET")
	program := gorillamux.New(router)
	return program.Proxy(req)
}

func main(){
	lambda.Start(Handler)
}