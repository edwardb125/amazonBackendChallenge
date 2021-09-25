# Amazon backend challenge

## API
url : 

## Project description
main.go file set router with gorilla mux and call SetDevice for post method and GetDevice for Get methode from controllers package.

in GetDevice connect to dynamoDB and search id and in SetDevice get data from request and change it to object and validate it after that connect to dynamoDb for sending data to it.

in CreateDynamoDB have some funtion thats duty is connect to dybamoDb that use in controllers functions.

in service file createDevice create device in table and GetDevice have some struct and get device from table.

in dynamoDb file we have some interface about PutItem and GetItem. There is Device struct that should be stored and send in deviceInformation in model folder.

## workflow
In this file some set like serverless deploy and environment variable and etc done to github can bulid project on vm machaine. workflow file is in .github folder.

## serverless
In serverless.yml file some set done like stage, region, ACCESS_TOKEN, SECRET_KEY and TABLE_NAME, also some dynamoDB properties set here like AttributeDefinitions, KeySchema.
ReadCapacityUnits and WriteCapacityUnits set in this part. 

## Environment variable
for test files in pc SECRET_KEY and ACCESS_TOKEN should be set in cmd to can access to AWS services.

## Test
Each test file is in same folder with source file. 2 test file is in controllers file that test getDevice and setDevice file. 2 another test is about service and they are in service folder, this test files test createDevice and getDevice.
