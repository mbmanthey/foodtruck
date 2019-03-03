package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	pb "github.com/mbmanthey/foodtruck/proto"
)

const (
	tableName = "foodtruck"
)

type TruckRepository struct {
	session *session.Session
}

//Repository is an interface of the functions for managing food trucks.
type Repository interface {
	Create(*pb.Truck) error
	GetAll() ([]*pb.Truck, error)
	Get(id string) (*pb.Truck, error)
	Delete(id string) error
	DeleteAll() error
}

//Create adds a truck to the mongodb.
func (repo *TruckRepository) Create(truck *pb.Truck) error {
	svc := dynamodb.New(repo.session)
	av, err := dynamodbattribute.MarshalMap(truck)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}
	return nil
}

//GetAll will retrieve all trucks from mongodb.
func (repo *TruckRepository) GetAll() ([]*pb.Truck, error) {
	var trucks []*pb.Truck
	svc := dynamodb.New(repo.session)
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
	})
	err = dynamodbattribute.UnmarshalMap(result.Item, &trucks)
	return trucks, err
}

//Get will retrieve a single truck using an ID.
func (repo *TruckRepository) Get(id string) (*pb.Truck, error) {
	var truck *pb.Truck
	svc := dynamodb.New(repo.session)
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(id),
			},
		},
	})
	err = dynamodbattribute.UnmarshalMap(result.Item, &truck)
	return truck, err
}

func (repo *TruckRepository) Delete(id string) error {
	svc := dynamodb.New(repo.session)
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(id),
			},
		},
	}
	_, err := svc.DeleteItem(input)
	return err
}

func (repo *TruckRepository) DeleteAll() error {
	svc := dynamodb.New(repo.session)
	_, err := svc.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		fmt.Println("Got error deleting table:")
		fmt.Println(err.Error())
	}
	tableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Timestamp"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Longitude"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Latitude"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err = svc.CreateTable(tableInput)
	return err
}
