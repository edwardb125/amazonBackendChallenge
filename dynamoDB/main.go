package dynamoDB

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DeviceDynamoDB interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput,error)
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput,error)
}

