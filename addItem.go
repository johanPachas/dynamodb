package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
)

type Item struct {
	Identificador int
	Nombre        string
	Apellido      string
	Curso_id      string
	Edad          int
}

func AddDBItem(item Item, svc *dynamodb.DynamoDB) {
	av, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Estudiante"),
	}
	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Error al llamar putItem: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Item agregado satisfactoriamente")
}
