package repository

import (
	"errors"
	"fmt"

	app "example.com/modart/app"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

type Database struct {
	Client                          *dynamodb.DynamoDB
	UserTablename, ArticleTablename string
}

func InitDatabase() app.AppService {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Database{
		client:           dynamodb.New(sess),
		UserTablename:    "Users",
		ArticleTablename: "Articles",
	}
}
func (db *Database) LoginAuthor(email string) (*app.Author, error) {
	return &app.Author{}, nil
}
func (db *Database) CreateAuthor(author *app.Author) (*app.Author, error) {
	author.Id = uuid.New().String()
	entityParsed, err := dynamodbattribute.MarshalMap(author)
	if err != nil {
		return &app.Author{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.UserTablename),
	}

	_, err = db.Client.PutItem(input)
	if err != nil {
		return &app.Author{}, err
	}

	return author, nil
}
func (db *Database) ReadAuthor(id string) (*app.Author, error) {
	result, err := db.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.UserTablename),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return &app.Author{}, err
	}
	if result.Item == nil {
		msg := fmt.Sprintf("Author with id [ %s ] not found", id)
		return &app.Author{}, errors.New(msg)
	}
	var author app.Author
	err = dynamodbattribute.UnmarshalMap(result.Item, &author)
	if err != nil {
		return &app.Author{}, err
	}

	return &author, nil
}
func (db *Database) ReadAuthors() ([]*app.Author, error) {
	authors := []*app.Author{}
	filt := expression.Name("Id").AttributeNotExists()
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("firstName"),
		expression.Name("lastName"),
		expression.Name("email"),
		expression.Name("articles"),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return []*app.Author{}, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(db.UserTablename),
	}
	result, err := db.Client.Scan(params)

	if err != nil {

		return []*app.Author{}, err
	}

	for _, item := range result.Items {
		var author app.Author
		err = dynamodbattribute.UnmarshalMap(item, &author)
		authors = append(authors, &author)

	}

	return authors, nil
}
func (db *Database) UpdateAuthor(author *app.Author) (*app.Author, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(author)
	if err != nil {
		return &app.Author{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.UserTablename),
	}

	_, err = db.Client.PutItem(input)
	if err != nil {
		return &app.Author{}, err
	}

	return author, nil
}
func (db *Database) DeleteAuthor(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(db.UserTablename),
	}

	res, err := db.Client.DeleteItem(input)
	if res == nil {
		return errors.New(fmt.Sprintf("No author to delete: %s", err))
	}
	if err != nil {
		return errors.New(fmt.Sprintf("Got error calling DeleteItem: %s", err))
	}
	return nil
}


