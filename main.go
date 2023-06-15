package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"fmt"
)

func main() {
	// crear una sesion
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		fmt.Println("Error al iniciar sesión")
	}

	//Crear cliente dynamoDB
	svc := dynamodb.New(sess)

	//Agregar un elemento
	item := Item{Identificador: 3, Nombre: "bbbbb", Apellido: "cccc", Curso_id: "curso1", Edad: 14}
	AddDBItem(item, svc)
}
