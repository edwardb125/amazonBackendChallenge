package controllers

import (
	_ "amazonBackendChallenge/deviceService"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"net/http"
	"os"
	_ "testing"
)

type Error struct{
	Message string `json:"message"`
}

func CreateError(w http.ResponseWriter, err string, status int){
	w.WriteHeader(status)
	result, _ := json.Marshal(Error{
		Message: err,
	})
	_, _ = w.Write(result)
}


func GetDynamoDB() (*dynamodb.DynamoDB, error){
	region := os.Getenv("AWS_REGION")
	secretKey := os.Getenv("SECRET_KEY")
	accessToken := os.Getenv("ACCESS_TOKEN")
	credential := credentials.NewStaticCredentials(accessToken, secretKey, "")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credential,
	})

	if err != nil{
		log.Println(err)
		return &dynamodb.DynamoDB{}, err
	}
	return dynamodb.New(awsSession), nil
}

//func CreateItem(t *testing.T, item models.Device){
//	db, err := GetDynamoDB()
//	if err != nil{
//		t.Fatal("error occurred while connecting to dynamodb")
//	}
//	err = deviceService.NewCreateService(db).CreateDevice(item)
//	if err != nil {
//		t.Fatal("error occurred while device creating")
//	}
//}

//func DeleteItem(t *testing.T, id string) {
//	db, err := GetDynamoDB()
//	if err != nil{
//		t.Fatal("error occurred while connecting to dynamodb")
//	}
//	deleteItemInput := &dynamodb.DeleteItemInput{
//		TableName: aws.String(os.Getenv("TABLE_NAME")),
//		KEY: map[string]*dynamodb.AttributeValue{
//			"id": &dynamodb.AttributeValue{
//				S: aws.String(id),
//			},
//		},
//	}
//	_, err = db.DeleteItem(deleteItemInput)
//	if err != nil{
//		t.Fatal("error occurred while deleting item from dynamodb")
//	}
//}