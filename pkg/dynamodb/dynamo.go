package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	// DefaultPoolSize is the default pool size
	_region    = "ap-southeast-1"
	_endpoint  = "http://localhost:8000" // Default endpoint for DynamoDB Local
	_tableName = "default_table"         // Default table name
)

type dynamo struct {
	db            *dynamodb.DynamoDB
	Configuration *aws.Config
}

func NewDynamoDB(endpoint string) (*dynamo, error) {
	c := &dynamo{}
	c.Configuration = &aws.Config{
		Region:   aws.String(_region),
		Endpoint: aws.String(endpoint),
	}

	sess, err := session.NewSession(c.Configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	dbClient := dynamodb.New(sess)
	if dbClient == nil {
		return nil, fmt.Errorf("failed to create DynamoDB client")
	}

	return &dynamo{
		db: dbClient,
	}, nil
}

func (d *dynamo) WithRegion(region string) *dynamo {
	d.Configuration.Region = aws.String(region)
	return d
}

func (c *dynamo) Configure(opts ...Option) dynamo {
	for _, opt := range opts {
		opt(c)
	}

	return *c
}
func (d *dynamo) GetDB() *dynamodb.DynamoDB {
	return d.db
}
