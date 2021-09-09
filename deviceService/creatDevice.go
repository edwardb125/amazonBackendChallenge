package deviceService

import (
	"amazonBackendChallenge/dynamoDB"
	"amazonBackendChallenge/models"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

type CreateCore struct{
	db dynamoDB.DeviceDynamoDB
}

func NewCreateService(db dynamoDB.DeviceDynamoDB) *CreateCore {
	return &CreateCore{
		db: db,
	}
}

func (d *CreateCore) CreateDevice(entity models.Device) error {
	device, _ := dynamodbattribute.MarshalMap(entity)
	input := &dynamodb.PutItemInput{
		Item:      device,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	_, err := d.db.PutItem(input)
	if err != nil {
		log.Println(err)
		return errors.New("server error")
	}
	return nil
}

