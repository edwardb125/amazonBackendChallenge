package service

import (
	"amazonBackendChallenge/mocks"
	"amazonBackendChallenge/models"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)
func TestGetDevice(t *testing.T) {

	item:=map[string]*dynamodb.AttributeValue{"id":&dynamodb.AttributeValue{S: aws.String("1")}}

	tests := []struct {
		name           string
		item           map[string]*dynamodb.AttributeValue
		getItemError   error
		errorExpected  error
		outputExpected models.Device
	}{
		{name: "device not found",errorExpected: errors.New("device not found")},
		{name: "get some errore in getItem",getItemError: errors.New("err"),errorExpected: errors.New("server error")},
		{name: "well done",item: item,outputExpected: models.Device{Id: "1"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB := new(mocks.DeviceDynamoDB)
			mockDB.On("GetItem", mock.Anything).Return(&dynamodb.GetItemOutput{
				Item: test.item,
			}, test.getItemError)
			service := &GetCore{
				Db: mockDB,
			}

			output, err := service.GetDevice("")
			// check that is has error or no
			if err == nil {
				assert.Nil(t, test.errorExpected)
			} else {
				assert.Error(t, test.errorExpected, err.Error())
			}
			assert.Equal(t, test.outputExpected, output)

		})
	}
}