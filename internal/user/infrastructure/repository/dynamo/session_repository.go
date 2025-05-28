package dynamo

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	domain "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
)

func SessionRepository(db *dynamodb.DynamoDB) *DynamoDBRepository {
	_tableName := "sessions"
	_, err := db.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(_tableName),
	})
	if err != nil {
		log.Printf("Table does not exist. Attempting to create...")
		// Create the table if it doesn't exist
		_, err = db.CreateTable(&dynamodb.CreateTableInput{
			TableName: aws.String(_tableName),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("expires_at"),
					AttributeType: aws.String("N"), // N for Number (Unix timestamp is a number)
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("expires_at"),
					KeyType:       aws.String("RANGE"), // Sort key
				},
			},

			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to create session table")
		}
		log.Printf("Table '%s' created successfully.", _tableName)
	} else {
		log.Printf("Table '%s' already exists.", _tableName)
	}
	return New(db, _tableName)
}

// GetSession retrieves a session by its ID from the DynamoDB table.
func (r *DynamoDBRepository) GetSession(ctx context.Context, sessionID string) (*domain.Session, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(sessionID),
			},
		},
	}

	result, err := r.db.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("session not found")
	}

	var session domain.Session
	if err := dynamodbattribute.UnmarshalMap(result.Item, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

// CreateSession creates a new session in the DynamoDB table.
func (r *DynamoDBRepository) CreateSession(ctx context.Context, session *domain.Session) (string, error) {
	av, err := dynamodbattribute.MarshalMap(session)
	if err != nil {
		return "", err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.db.PutItemWithContext(ctx, input)
	if err != nil {
		return "", err
	}

	return session.ID, nil
}

func (r *DynamoDBRepository) DeleteSession(ctx context.Context, sessionID string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(sessionID),
			},
		},
	}

	_, err := r.db.DeleteItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r *DynamoDBRepository) UpdateSession(ctx context.Context, session *domain.Session) error {
	av, err := dynamodbattribute.MarshalMap(session)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.db.PutItemWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r *DynamoDBRepository) ListSessions(ctx context.Context) ([]*domain.Session, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	result, err := r.db.ScanWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	var sessions []*domain.Session
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *DynamoDBRepository) GetSessionByUserID(ctx context.Context, username string) ([]*domain.Session, error) {
	input := &dynamodb.QueryInput{
		TableName: aws.String(r.tableName),
		IndexName: aws.String("user_id-index"), // Assuming you have a GSI on user_id
		KeyConditions: map[string]*dynamodb.Condition{
			"username": {
				AttributeValueList: []*dynamodb.AttributeValue{
					{S: aws.String(username)},
				},
				ComparisonOperator: aws.String("EQ"),
			},
		},
	}

	result, err := r.db.QueryWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	var sessions []*domain.Session
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *DynamoDBRepository) BlockSession(ctx context.Context, sessionID string) (bool, error) {
	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(sessionID)},
		},
		UpdateExpression:    aws.String("SET is_blocked = :is_blocked"),
		ConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":is_blocked": {BOOL: aws.Bool(true)},
			":id":         {S: aws.String(sessionID)},
		},
		ReturnValues: aws.String("UPDATED_NEW"), // Optional: what to return after update
	}

	_, err := r.db.UpdateItem(updateItemInput)
	if err != nil {
		log.Error().Err(err).Msg("failed to update password in DynamoDB")
		return false, err
	}
	return true, nil

}

func (r *DynamoDBRepository) UnblockSession(ctx context.Context, sessionID string) (bool, error) {
	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(sessionID)},
		},
		UpdateExpression:    aws.String("SET is_blocked = :is_blocked"),
		ConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":is_blocked": {BOOL: aws.Bool(false)},
			":id":         {S: aws.String(sessionID)},
		},
		ReturnValues: aws.String("UPDATED_NEW"), // Optional: what to return after update
	}

	_, err := r.db.UpdateItem(updateItemInput)
	if err != nil {
		log.Error().Err(err).Msg("failed to update password in DynamoDB")
		return false, err
	}
	return true, nil
}

// GenerateSession creates a new session with a unique ID and an expiration time.
func (r *DynamoDBRepository) GenerateSession(ctx context.Context, username string, token string, duration time.Duration) (*domain.Session, error) {
	session := &domain.Session{
		ID:        uuid.New().String(),
		Username:  username,
		ExpiresAt: time.Now().Add(duration).Unix(),
		IsBlocked: false,
		Token:     token,
	}

	_, err := r.CreateSession(ctx, session)
	if err != nil {
		log.Error().Err(err).Msg("failed to create session")
		return nil, err
	}

	return session, nil
}
