package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

// NewDynamoDBUserRepository creates a new instance of DynamoDBUserRepository.
func New(db *dynamodb.DynamoDB, tableName string) *DynamoDBRepository {
	return &DynamoDBRepository{db: db, tableName: tableName}
}
