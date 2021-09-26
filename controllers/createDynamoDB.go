package controllers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"testing"
)

type Error struct{
	Message string `json:"message"`
}

func DeleteDeviceId(t *testing.T, id string) {
	db, err := ConnectDynamoDB()
	if err != nil{
		t.Fatal("error occurred while connecting to dynamodb")
	}
	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
	}
	_, err = db.DeleteItem(deleteItemInput)

	if err != nil{
		t.Fatal("error occurred while deleting item from dynamodb")
	}
}

func ConnectDynamoDB() (*dynamodb.DynamoDB, error){
	// Set environment Token
	accessToken := os.Getenv("ACCESS")
	secretKey := os.Getenv("SECRET")
	region := os.Getenv("AWS_REGION")
	temp := credentials.NewStaticCredentials(accessToken, secretKey, "")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: temp,
	})

	if err != nil{
		return &dynamodb.DynamoDB{}, err
	}
	return dynamodb.New(awsSession), nil
}


