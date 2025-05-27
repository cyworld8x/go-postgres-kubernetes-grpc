package dynamodb

import "github.com/aws/aws-sdk-go/service/dynamodb"

type IDynamoDB interface {
	GetDB() *dynamodb.DynamoDB
	WithRegion(region string) *dynamodb.DynamoDB
}
