package dynamo

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	domain "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/domain"
)

const (
	_tableName = "users"
)

// NewDynamoDBRepository creates a new instance of DynamoDBRepository.
func NewUserRepository(db *dynamodb.DynamoDB) *DynamoDBRepository {
	// Check if the table exists
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
					AttributeName: aws.String("username"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("username"),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to create table")
		}
		log.Printf("Table '%s' created successfully.", _tableName)
	} else {
		log.Printf("Table '%s' already exists.", _tableName)
	}

	return New(db, _tableName)
}

// CreateUser implements the UserRepository interface for DynamoDB.
func (r *DynamoDBRepository) CreateUser(ctx context.Context, user *domain.User) error {
	av, err := dynamodbattribute.MarshalMap(user)
	av["ID"] = &dynamodb.AttributeValue{S: aws.String(user.ID.String())}
	if err != nil {
		return fmt.Errorf("failed to marshal user for DynamoDB: %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.tableName),
	}

	_, err = r.db.PutItemWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put item in DynamoDB: %w", err)
	}
	return nil
}

// GetLogin implements the UserRepository interface for DynamoDB.
func (r *DynamoDBRepository) GetLogin(ctx context.Context, username string) (*domain.User, error) {
	input := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("#st = :status AND username = :username"),
		ExpressionAttributeNames: map[string]*string{
			"#st": aws.String("status"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":username": {S: aws.String(username)},
			":status":   {BOOL: aws.Bool(true)},
		},
	}

	result, err := r.db.ScanWithContext(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get item from DynamoDB")
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, errors.New("user does not exist") // Specific error for not found
	}

	var rawUser domain.RawUser
	if err := dynamodbattribute.UnmarshalMap(result.Items[0], &rawUser); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal user from DynamoDB")
		return nil, err
	}

	id, err := uuid.Parse(rawUser.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse UUID")
	}

	return &domain.User{
		ID:          id,
		Username:    rawUser.Username,
		Password:    rawUser.Password,
		Code:        rawUser.Code,
		Email:       rawUser.Email,
		DisplayName: rawUser.DisplayName,
		Role:        rawUser.Role,
		Created:     rawUser.Created,
		Updated:     rawUser.Updated,
		Status:      rawUser.Status,
	}, nil
}

// GetLogin implements the UserRepository interface for DynamoDB.
func (r *DynamoDBRepository) GetUser(ctx context.Context, id string) (*domain.User, error) {
	input := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("#st = :status AND ID = :id"),
		ExpressionAttributeNames: map[string]*string{
			"#st": aws.String("status"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id":     {S: aws.String(id)},
			":status": {BOOL: aws.Bool(true)},
		},
	}

	result, err := r.db.ScanWithContext(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get item from DynamoDB")
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, errors.New("user dosn't exist") // Specific error for not found
	}

	var rawUser domain.RawUser
	if err := dynamodbattribute.UnmarshalMap(result.Items[0], &rawUser); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal user from DynamoDB")
		return nil, err
	}

	userId, err := uuid.Parse(rawUser.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse UUID")
	}

	return &domain.User{
		ID:          userId,
		Username:    rawUser.Username,
		Code:        rawUser.Code,
		Password:    rawUser.Password,
		Email:       rawUser.Email,
		DisplayName: rawUser.DisplayName,
		Role:        rawUser.Role,
		Created:     rawUser.Created,
		Updated:     rawUser.Updated,
		Status:      rawUser.Status,
	}, nil
}

// UpdateUser implements the UserRepository interface for DynamoDB.
func (r *DynamoDBRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	// PutItem acts as an upsert. If the item exists, it updates it.
	// For a true "update" operation (e.g., only updating specific fields),
	// you might use UpdateItem API with ExpressionAttributeValues and UpdateExpression.
	// For this example, PutItem is sufficient for full replacement/upsert.
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal user for DynamoDB update")
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.tableName),
	}

	_, err = r.db.PutItemWithContext(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("failed to put item in DynamoDB for update")
		// Return a more specific error if needed
		return err
	}
	return nil
}

func (r *DynamoDBRepository) ChangePassword(ctx context.Context, id string, username string, password string) error {

	updateItemInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(username)},
		},
		UpdateExpression:    aws.String("SET password = :password"),
		ConditionExpression: aws.String("ID = :id AND username = :username"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":password": {S: aws.String(password)},
			":id":       {S: aws.String(id)},
			":username": {S: aws.String(username)},
		},
		ReturnValues: aws.String("UPDATED_NEW"), // Optional: what to return after update
	}

	_, err := r.db.UpdateItem(updateItemInput)
	if err != nil {
		log.Error().Err(err).Msg("failed to update password in DynamoDB")
		return err
	}

	return nil

}

// DeleteUser implements the UserRepository interface for DynamoDB.
func (r *DynamoDBRepository) DeleteUser(ctx context.Context, id string) error {

	user, err := r.GetUser(ctx, id)

	if err != nil {
		log.Error().Err(err).Msg("failed to get user for deletion")
		return err
	}

	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(user.Username)},
		},
		ConditionExpression: aws.String("ID = :id AND username = :username"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id":       {S: aws.String(id)},
			":username": {S: aws.String(user.Username)},
		},
	}

	_, err = r.db.DeleteItemWithContext(ctx, deleteItemInput)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete user from DynamoDB")
		return err
	}
	return nil
}
